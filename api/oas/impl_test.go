package oas

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/woocoos/knockout-go/api/fs"
	"github.com/woocoos/knockout-go/api/fs/alioss"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/msgcenter/ent/msgalert"
	"github.com/woocoos/msgcenter/ent/msginternal"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/notify/webhook"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service"
	"github.com/woocoos/msgcenter/service/provider/mem"
	"github.com/woocoos/msgcenter/test/maildev"
	"github.com/woocoos/msgcenter/test/testsuite"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

func init() {
	fs.RegisterS3Provider(fs.KindAliOSS, alioss.BuildProvider)
}

// ServiceSuite is the service test suite
type serviceSuite struct {
	testsuite.BaseSuite

	server    *ServerImpl
	shutdowns []func()
	maildev   *maildev.MailDev

	webhook        *httptest.Server
	webhookHandler http.Handler
}

// TestServiceSuite runs the service test suite
func TestServiceSuite(t *testing.T) {
	s := &serviceSuite{
		maildev: maildev.DefaultServer(),
	}
	s.DSN = "file:msgcenter?mode=memory&cache=shared&_fk=1"
	s.DriverName = "sqlite3"
	s.webhook = httptest.NewUnstartedServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/token" {
				w.Header().Set("Content-Type", "application/json")
				d, err := json.Marshal(map[string]string{
					"access_token": "90d64460d14870c08c81352a05dedd3465940a7c",
					"expires_in":   "7200", // defaultExpiryDelta = 10 * time.Second, so set 11 seconds and sleep 1 second
					"scope":        "user",
					"token_type":   "bearer",
				})
				require.NoError(t, err)
				w.Write(d)
				return
			} else if r.URL.Path == "/webhook" {
				if s.webhookHandler != nil {
					s.webhookHandler.ServeHTTP(w, r)
				}
			} else if r.URL.Path == "/graphql/query" {
				w.Header().Set("Content-Type", "application/json")
				body, _ := io.ReadAll(r.Body)
				if strings.Contains(string(body), "fileIdentities") {
					d, err := json.Marshal(map[string]any{
						"data": map[string]any{
							"fileIdentitiesForApp": []map[string]any{
								{
									"id": 1, "tenantID": 1,
									"accessKeyID": "test-ak", "accessKeySecret": "test-sk",
									"roleArn": "", "policy": "", "durationSeconds": 3600,
									"isDefault": true,
									"source": map[string]any{
										"id": 1, "kind": "aliOSS",
										"endpoint":          "https://oss-cn-hangzhou.aliyuncs.com",
										"endpointImmutable": false,
										"stsEndpoint":       "https://sts.cn-hangzhou.aliyuncs.com",
										"region":            "cn-hangzhou",
										"bucket":            "test-bucket",
										"bucketURL":         "https://test-bucket.oss-cn-hangzhou.aliyuncs.com",
									},
								},
							},
						},
					})
					require.NoError(t, err)
					w.Write(d)
					return
				}
				d, err := json.Marshal(map[string]string{})
				require.NoError(t, err)
				w.Write(d)
				return
			} else if r.URL.Path == "/org/domain" {
				w.Header().Set("Content-Type", "application/json")
				d, err := json.Marshal(map[string]any{
					"id":              2,
					"local_currency":  "HKD",
					"name":            "组织1",
					"parent_id":       1,
					"parent_name":     "组织2",
					"parent_currency": "HKD",
				})
				require.NoError(t, err)
				w.Write(d)
				return
			}
			return
		}))
	var err error
	s.webhook.Listener, err = net.Listen("tcp", "127.0.0.1:5001")
	require.NoError(t, err)
	s.webhook.Start()
	defer s.webhook.Close()
	suite.Run(t, s)
}

// SetupSuite sets up the test suite
func (s *serviceSuite) SetupSuite() {
	err := s.BaseSuite.Setup()
	s.Require().NoError(err)
	s.Require().NoError(s.initData())
	s.AlertManager, err = service.NewAlertManager(s.App, service.WithClient(s.Client))
	s.Require().NoError(err)

	s.server, err = NewServer(s.App, s.AlertManager, nil)
	s.Require().NoError(err)

	s.AlertManager.Coordinator.ReloadHooks(func(c *profile.Config) error {
		s.AlertManager.Coordinator.Template.ExternalURL, err = url.Parse("http://localhost:9093")
		s.Require().NoError(err)
		s.Require().NoError(s.AlertManager.Start(s.AlertManager.Coordinator, c))

		s.server.Update(c, func(labels label.LabelSet) {
			ctx := context.Background()
			s.AlertManager.Inhibitor.Load().Mutes(ctx, labels)
			s.AlertManager.Silencer.Mutes(ctx, labels)
		})

		return nil
	})

	err = s.AlertManager.Coordinator.Reload()
	s.Require().NoError(err)
	alerts := s.AlertManager.Alerts.(*mem.Alerts)
	go alerts.Start(context.Background())
	s.shutdowns = append(s.shutdowns, func() {
		s.AlertManager.Stop()
		alerts.Stop(context.Background())
	})
}

// 在此添加特殊的用例数据
func (s *serviceSuite) initData() error {
	ctx := s.NewTestCtx()

	// Create a user-level template for user 1 on AlterPassword event.
	// This template has a distinct subject to verify user-level template priority.
	s.Client.MsgTemplate.Create().
		SetMsgTypeID(1).SetEventID(1).SetTenantID(1).SetUserID(1).
		SetName("UserCustomAlterPassword").SetCreatedBy(1).
		SetStatus(typex.SimpleStatusActive).
		SetFormat(msgtemplate.FormatTxt).
		SetReceiverType(profile.ReceiverEmail).
		SetTo(`{{ template "email.to" . }}`).
		SetSubject(`用户定制模板测试`).
		SetBody(`{{ template "1.alterpwd.txt" . }}`).
		SaveX(ctx)

	return nil
}

// TearDownSuite tears down the test suite
func (s *serviceSuite) TearDownSuite() {
	for _, shutdown := range s.shutdowns {
		shutdown()
	}
}

// test postalerts
func (s *serviceSuite) TestPostAlerts() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"alertname":         "AlterPassword",
					label.TenantLabel:   "1",
					label.ToUserIDLabel: "3",
				},
			},
			Annotations: map[string]string{
				"summary": "summary",
				"text":    "text",
			},
			EndsAt:   new(time.Now().Add(time.Hour)),
			StartsAt: new(time.Now()),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))
	time.Sleep(time.Second * 2)
	mail, err := s.maildev.GetLastEmail()
	s.Require().NoError(err)
	s.Require().NotNil(mail)
	s.Require().Equal("nobody@localhost", mail.To[0]["Address"])
}

func (s *serviceSuite) TestPostAlerts_TraceIDHeader() {
	countBefore, err := s.maildev.MessageCount()
	s.Require().NoError(err)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"alertname":         "AlterPassword",
					label.TenantLabel:   "1",
					label.ToUserIDLabel: "1",
				},
			},
			Annotations: map[string]string{
				"summary": "trace id test",
				"text":    "text",
			},
			EndsAt:   new(time.Now().Add(time.Hour)),
			StartsAt: new(time.Now()),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))

	// 等待邮件到达
	var mail *maildev.MailDevEmail
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		count, err := s.maildev.MessageCount()
		s.Require().NoError(err)
		if count > countBefore {
			mail, err = s.maildev.GetEmailAt(0)
			s.Require().NoError(err)
			break
		}
	}
	s.Require().NotNil(mail, "email should arrive within 10 seconds")

	// 获取完整邮件, MessageID 即为 trace ID (同一 rawID)
	msg, err := s.maildev.GetMessage(mail.ID)
	s.Require().NoError(err)
	s.NotEmpty(msg.MessageID, "Message-ID should be present")
	traceID := msg.MessageID

	// 验证 Nlog 中存储了相同的 message_id
	// Mailpit 返回的 MessageID 格式为 "rawID@hostname", Nlog 存的是 "rawID"
	nlogs, err := s.Client.Nlog.Query().All(schemax.SkipTenantPrivacy(context.Background()))
	s.Require().NoError(err)
	s.NotEmpty(nlogs, "should have at least one Nlog")
	var found bool
	for _, nl := range nlogs {
		if nl.MessageID != "" && strings.HasPrefix(traceID, nl.MessageID) {
			found = true
			break
		}
	}
	s.True(found, "Nlog message_id should be a prefix of the email Message-ID")
}

// TestPostAlertsWithDynamicAttachments tests email notification with dynamic attachments
// from alert annotations, including both HTTP URL and local file path.
func (s *serviceSuite) TestPostAlertsWithDynamicAttachments() {
	// Record message count before sending to locate our email precisely.
	countBefore, err := s.maildev.MessageCount()
	s.Require().NoError(err)

	// Create a test file to serve as HTTP attachment
	tmpFile, err := os.CreateTemp("", "dynamic-attachment-*.txt")
	s.Require().NoError(err)
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.WriteString("dynamic attachment content")
	s.Require().NoError(err)
	tmpFile.Close()

	// Start HTTP test server to serve the attachment file
	attServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", `attachment; filename="dynamic.txt"`)
		w.Header().Set("Content-Type", "text/plain")
		http.ServeFile(w, r, tmpFile.Name())
	}))
	defer attServer.Close()

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"alertname":       "AlterPassword",
					label.TenantLabel: "1",
					// Use unique user to create a distinct group, avoiding group_interval delay.
					label.ToUserIDLabel: "dynatt",
				},
			},
			Annotations: map[string]string{
				"summary":                         "dynamic attachment test",
				"to":                              "alerts@example.com",
				alert.DynamicAttachmentAnnotation: attServer.URL + "/dynamic.txt",
			},
			EndsAt:   new(time.Now().Add(time.Hour)),
			StartsAt: new(time.Now()),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))

	// Poll for the email to arrive: wait for message count to increase.
	var mail *maildev.MailDevEmail
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		count, err := s.maildev.MessageCount()
		s.Require().NoError(err)
		if count > countBefore {
			// New email arrived, it's at index 0 (newest-first order).
			mail, err = s.maildev.GetEmailAt(0)
			s.Require().NoError(err)
			break
		}
	}
	s.Require().NotNil(mail, "email should arrive within 10 seconds")
	s.Require().Equal("alerts@example.com", mail.To[0]["Address"])
	s.Require().Greater(mail.Attachments, 0, "email should have at least 1 attachment")

	// Verify attachment details via full message.
	msg, err := s.maildev.GetMessage(mail.ID)
	s.Require().NoError(err)
	s.Require().Len(msg.Attachments, 1)
	s.Require().Equal("dynamic.txt", msg.Attachments[0].FileName)
	s.Require().Equal("text/plain", msg.Attachments[0].ContentType)
}

// TestPostAlertsWithDynamicAttachments_OSSMount tests that OSS URLs in annotations
// are resolved to local mount paths at the API stage, so the email notifier
// attaches the file directly from the mounted filesystem.
func (s *serviceSuite) TestPostAlertsWithDynamicAttachments_OSSMount() {
	// Record message count before sending to locate our email precisely.
	countBefore, err := s.maildev.MessageCount()
	s.Require().NoError(err)

	// Verify KOSdk was initialized with the mock provider.
	s.Require().NotNil(s.server.coordinator.KOSdk, "KOSdk must be initialized")
	s.Require().NotEmpty(s.server.coordinator.MountPaths, "mountPaths must be configured")

	// Create a local file at the expected mount path.
	mountDir := "tmp/oss-mount/test-bucket"
	s.Require().NoError(os.MkdirAll(mountDir, 0o755))
	localFile, err := filepath.Abs(mountDir + "/test-attachment.txt")
	s.Require().NoError(err)
	s.Require().NoError(os.WriteFile(localFile, []byte("oss mount test content"), 0o644))
	defer os.Remove(localFile)

	// The OSS URL should match the mock provider's BucketUrl.
	ossURL := "https://test-bucket.oss-cn-hangzhou.aliyuncs.com/test-attachment.txt"

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"alertname":       "AlterPassword",
					label.TenantLabel: "1",
					// Use unique user to create a distinct group, avoiding group_interval delay.
					label.ToUserIDLabel: "ossmount",
				},
			},
			Annotations: map[string]string{
				"to":                              "alerts@example.com",
				"summary":                         "oss mount attachment test",
				alert.DynamicAttachmentAnnotation: ossURL,
			},
			EndsAt:   new(time.Now().Add(time.Hour)),
			StartsAt: new(time.Now()),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))

	// Poll for the email to arrive: wait for message count to increase.
	var mail *maildev.MailDevEmail
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		count, err := s.maildev.MessageCount()
		s.Require().NoError(err)
		if count > countBefore {
			// New email arrived, it's at index 0 (newest-first order).
			mail, err = s.maildev.GetEmailAt(0)
			s.Require().NoError(err)
			break
		}
	}
	s.Require().NotNil(mail, "email should arrive within 10 seconds")

	// Verify the email was sent with the local file as attachment.
	s.Require().Equal("alerts@example.com", mail.To[0]["Address"])
	s.Require().Greater(mail.Attachments, 0, "email should have at least 1 attachment")

	msg, err := s.maildev.GetMessage(mail.ID)
	s.Require().NoError(err)
	s.Require().Len(msg.Attachments, 1)
	s.Require().Equal("test-attachment.txt", msg.Attachments[0].FileName)
}

// TestPostAlertsWithParams
func (s *serviceSuite) TestPostAlertsWithParams() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"alertname":       "AlterPassword",
					label.TenantLabel: "1",
				},
			},
			Annotations: map[string]string{
				"to":       "alerts@example.com",
				"nickname": "test",
			},
			EndsAt:   new(time.Now().Add(time.Hour)),
			StartsAt: new(time.Now()),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))
	time.Sleep(time.Second * 2)
	mail, err := s.maildev.GetLastEmail()
	s.Require().NoError(err)
	s.Require().NotNil(mail)
	s.Require().Equal("alerts@example.com", mail.To[0]["Address"])
}

// TestPostAlertsWithTenant tenant with custom template and attachment
func (s *serviceSuite) TestPostAlertsWithTenant() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	// rand a string
	to := fmt.Sprintf("%d@localhost", rand.Intn(10000000))
	// route: default
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"receiver":           "email|webhook",
					label.AlertNameLabel: "noSubscribe",
					label.TenantLabel:    "1",
				},
			},
			Annotations: map[string]string{
				"to":       to,
				"summary":  "test",
				"nickname": "woocoos",
			},
			EndsAt:   new(time.Now().Add(time.Second * 2)),
			StartsAt: new(time.Now()),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))
	time.Sleep(time.Second * 3)
	mail, err := s.maildev.GetLastEmail()
	s.Require().NoError(err)
	s.Require().NotNil(mail)
	s.Require().Equal(to, mail.To[0]["Address"])
}

// TestPostAlertsWithDefaultTpl tenant with custom template and attachment
func (s *serviceSuite) TestPostAlertsWithDefaultTpl() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	// rand a string
	to := fmt.Sprintf("%d@localhost", rand.Intn(10000000))
	// route: default
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"receiver":           "email",
					label.AlertNameLabel: "defaultparams",
					label.TenantLabel:    "1",
				},
			},
			Annotations: map[string]string{
				"to":        to,
				"nickname":  "test",
				"timestamp": strconv.Itoa(int(time.Now().Unix())),
			},
			EndsAt:   new(time.Now().Add(time.Hour)),
			StartsAt: new(time.Now()),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))
	time.Sleep(time.Second * 10)
	mail, err := s.maildev.GetLastEmail()
	s.Require().NoError(err)
	s.Require().NotNil(mail)
	s.Require().Equal(to, mail.To[0]["Address"])
}

func (s *serviceSuite) TestPostAlertsWithGroupBy() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	// rand a string
	to := fmt.Sprintf("%d@localhost", rand.Intn(10000000))
	// route: default
	req1 := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"receiver":           "email",
					label.AlertNameLabel: "MsgGroupBy",
					label.TenantLabel:    "1",
					"timestamp":          strconv.Itoa(int(time.Now().Unix())),
				},
			},
			Annotations: map[string]string{
				"to":       to,
				"nickname": "test1",
				"message":  "msg1",
			},
		},
	}
	fmt.Println("===start 1===")
	err := s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req1})
	s.Require().NoError(err)
	// sleep5秒，等待req1消息发送完成
	time.Sleep(time.Second * 5)
	fmt.Println("===start 2===")
	req2 := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"receiver":           "email",
					label.AlertNameLabel: "MsgGroupBy",
					label.TenantLabel:    "1",
					"timestamp":          strconv.Itoa(int(time.Now().Add(1 * time.Second).Unix())),
				},
			},
			Annotations: map[string]string{
				"to":       to,
				"nickname": "test1",
				"message":  "msg2",
			},
		},
	}
	err = s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req2})
	s.Require().NoError(err)
	// sleep5秒，等待req1消息发送完成
	time.Sleep(time.Second * 5)
	fmt.Println("===start 3===")
	req3 := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"receiver":           "email",
					label.AlertNameLabel: "MsgGroupBy",
					label.TenantLabel:    "1",
					"timestamp":          strconv.Itoa(int(time.Now().Add(2 * time.Second).Unix())),
				},
			},
			Annotations: map[string]string{
				"to":       to,
				"nickname": "test1",
				"message":  "msg3",
			},
		},
		{
			Alert: &Alert{
				Labels: map[string]string{
					"receiver":           "email",
					label.AlertNameLabel: "MsgGroupBy",
					label.TenantLabel:    "1",
					"timestamp":          strconv.Itoa(int(time.Now().Add(3 * time.Second).Unix())),
				},
			},
			Annotations: map[string]string{
				"to":       to,
				"nickname": "test1",
				"message":  "msg4",
			},
		},
	}
	err = s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req3})
	s.Require().NoError(err)
	time.Sleep(time.Second * 30)
	mail, err := s.maildev.GetLastEmail()
	s.Require().NoError(err)
	s.Require().NotNil(mail)
	s.Require().Equal(to, mail.To[0]["Address"])

	//二次测试
}

// TestPostAlertsWithTenant tenant with custom template and attachment
func (s *serviceSuite) TestUserSubscribe() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	// route: default
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"receiver":           "email|webhook",
					label.AlertNameLabel: testsuite.SubEventName,
					"app":                "1",
					"tenant":             "1",
				},
			},
			Annotations: map[string]string{
				"summary":  "test",
				"nickname": "woocoos",
				"to":       "test@test.com",
			},
			EndsAt:   new(time.Now().Add(time.Second * 5)),
			StartsAt: new(time.Now()),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))
	time.Sleep(time.Second * 3)
	mail, err := s.maildev.GetLastEmail()
	s.Require().NoError(err)
	s.Require().NotNil(mail)
	// user 1 or user2,but notify not keep order
	s.Require().Equal("订阅事件提醒", mail.Subject)

	ss, err := s.Client.MsgAlert.Query().Where(msgalert.TenantID(1), func(selector *sql.Selector) {
		selector.Where(sqljson.ValueEQ(msgalert.FieldLabels, testsuite.SubEventName, sqljson.Path("alertname")))
	}).All(schemax.SkipTenantPrivacy(context.Background()))
	s.Require().NoError(err)
	s.Require().Len(ss, 1)
}

func (s *serviceSuite) TestWebhook() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	var got webhook.Message
	s.webhookHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &got)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"event":    "app:approve",
					"receiver": "webhook",
					"skipSub":  "Y",
					"tenant":   "1000",
				},
			},
			Annotations: map[string]string{
				"summary": "webhook test",
				"mobile":  "8618359260323",
			},
			StartsAt: new(time.Now()),
			EndsAt:   new(time.Now().Add(time.Hour)),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))
	time.Sleep(time.Second * 3)
	s.Require().NotNil(got.Data)
	s.Require().Equal("webhook test", got.Data.CommonAnnotations["summary"])
}

func (s *serviceSuite) TestWebhook_CustomTpl_DingTalk() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	var got string
	s.webhookHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		got = string(body)
	})
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					label.AlertNameLabel: testsuite.WebhookEventName,
					"event":              "app:approve",
					"receiver":           "webhook",
					"tenant":             "1",
					"skipSub":            "Y",
					"severity":           "critical",
				},
			},
			Annotations: map[string]string{
				"summary": "webhook template test",
			},
			EndsAt:   new(time.Now().Add(time.Hour)),
			StartsAt: new(time.Now().Add(time.Second * 5)),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))
	time.Sleep(time.Second * 3)
	s.Require().Contains(got, "webhook template test")
}

// TestWebhook_Dingtalk sends a real DingTalk webhook notification to verify
// the DingTalk markdown message format.
func (s *serviceSuite) TestWebhook_Dingtalk() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	startsAt := time.Now()
	endsAt := time.Now().Add(time.Hour)
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					label.AlertNameLabel: testsuite.DingtalkEventName,
					"event":              "app:approve",
					"receiver":           "dingtalk",
					"tenant":             "1",
					"skipSub":            "Y",
					"severity":           "critical",
				},
			},
			Annotations: map[string]string{
				"to":          "tst@qq.com",
				"summary":     "钉钉webhook测试",
				"description": "这是一条测试告警消息",
			},
			StartsAt: &startsAt,
			EndsAt:   &endsAt,
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))
	// Wait for notification pipeline to process.
	time.Sleep(time.Second * 10)
}

func (s *serviceSuite) TestMessage() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"tenant":  "1",
					"event":   "app:message",
					"skipSub": "Y",
					"user":    "1,2",
				},
			},
			Annotations: map[string]string{
				"summary": "internal message test",
			},
			StartsAt: new(time.Now()),
			EndsAt:   new(time.Now().Add(time.Hour)),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))
	time.Sleep(time.Second * 2)
	// init test has insert 2 message
	mis, err := s.Client.MsgInternal.Query().Where(msginternal.IDGT(2)).
		All(schemax.SkipTenantPrivacy(context.Background()))
	s.Require().NoError(err)
	s.Len(mis, 1)
	mist, err := mis[0].MsgInternalTo(schemax.SkipTenantPrivacy(context.Background()))
	s.Require().NoError(err)
	s.Len(mist, 2)
}

// TestUserLevelTemplate verifies that user-level templates take priority over tenant-level templates.
func (s *serviceSuite) TestUserLevelTemplate() {
	countBefore, err := s.maildev.MessageCount()
	s.Require().NoError(err)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"alertname":       "AlterPassword",
					label.TenantLabel: "1",
					// User 1 has a custom user-level template.
					label.ToUserIDLabel: "1",
				},
			},
			Annotations: map[string]string{
				"summary": "user template test",
			},
			EndsAt:   new(time.Now().Add(time.Hour)),
			StartsAt: new(time.Now()),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))

	// Poll for the email to arrive.
	var mail *maildev.MailDevEmail
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		count, err := s.maildev.MessageCount()
		s.Require().NoError(err)
		if count > countBefore {
			mail, err = s.maildev.GetEmailAt(0)
			s.Require().NoError(err)
			break
		}
	}
	s.Require().NotNil(mail, "email should arrive within 10 seconds")
	// User-level template subject should be used.
	s.Require().Equal("用户定制模板测试", mail.Subject)
}

// TestTenantLevelTemplateFallback verifies that when no user-level template exists,
// the system falls back to the tenant-level template.
func (s *serviceSuite) TestTenantLevelTemplateFallback() {
	countBefore, err := s.maildev.MessageCount()
	s.Require().NoError(err)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := PostableAlerts{
		{
			Alert: &Alert{
				Labels: map[string]string{
					"alertname":       "AlterPassword",
					label.TenantLabel: "1",
					// User 2 has NO user-level template, should fall back to tenant-level.
					label.ToUserIDLabel: "2",
				},
			},
			Annotations: map[string]string{
				"summary": "tenant fallback test",
			},
			EndsAt:   new(time.Now().Add(time.Hour)),
			StartsAt: new(time.Now()),
		},
	}
	s.Require().NoError(s.server.PostAlerts(ctx, &PostAlertsRequest{PostableAlerts: req}))

	// Poll for the email to arrive.
	var mail *maildev.MailDevEmail
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		count, err := s.maildev.MessageCount()
		s.Require().NoError(err)
		if count > countBefore {
			mail, err = s.maildev.GetEmailAt(0)
			s.Require().NoError(err)
			break
		}
	}
	s.Require().NotNil(mail, "email should arrive within 10 seconds")
	// Tenant-level template subject contains "密码到期提醒".
	s.Require().Contains(mail.Subject, "密码到期提醒")
	// Should NOT be the user-level template subject.
	s.Require().NotEqual("用户定制模板测试", mail.Subject)
}

func (s *serviceSuite) TestPostSilence() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	req := PostableSilence{
		Silence: &Silence{
			Comment:   "test",
			CreatedBy: 1,
			StartsAt:  time.Now(),
			EndsAt:    time.Now().Add(time.Hour),
			Matchers: []*Matcher{
				{
					Name:  "alertname",
					Value: "test",
				},
			},
		},
	}
	res, err := s.server.PostSilences(ctx, &PostSilencesRequest{PostableSilence: req})
	s.Require().NoError(err)
	s.NotZero(res.SilenceID)
}

func (s *serviceSuite) TestUpdateNlog() {
	ctx := s.NewTestCtx()

	// 创建 MsgAlert 作为关联目标
	alertLabels := label.LabelSet{"alertname": "TestNlogErr"}
	msAlert := s.Client.MsgAlert.Create().
		SetTenantID(1).SetLabels(&alertLabels).
		SetFingerprint("test-nlog-err").SetState(alert.AlertFiring).
		SetStartsAt(time.Now()).SetEndsAt(time.Now().Add(time.Hour)).
		SaveX(ctx)

	// 创建两条相同 alert+receiverType 的 Nlog, 模拟多次通知发送
	nlog1 := s.Client.Nlog.Create().
		SetTenantID(1).SetReceiver("email").SetGroupKey("test-group").
		SetReceiverType(profile.ReceiverEmail).SetIdx(0).
		SetExpiresAt(time.Now().Add(24*time.Hour)).SetSendAt(time.Now()).
		AddAlertIDs(msAlert.ID).
		SaveX(ctx)

	nlog2 := s.Client.Nlog.Create().
		SetTenantID(1).SetReceiver("email").SetGroupKey("test-group").
		SetReceiverType(profile.ReceiverEmail).SetIdx(0).
		SetExpiresAt(time.Now().Add(24*time.Hour)).SetSendAt(time.Now()).
		AddAlertIDs(msAlert.ID).
		SaveX(ctx)

	gc, _ := gin.CreateTestContext(httptest.NewRecorder())
	gc.Request = httptest.NewRequest(http.MethodPost, "/nlogs/update", nil)
	gc.Request = gc.Request.WithContext(ctx)

	// FIFO 顺序: 最早的未处理 Nlog (nlog1) 应优先被更新
	resp, err := s.server.UpdateNlog(gc, &UpdateNlogRequest{
		NlogUpdate: NlogUpdate{
			AlertID:      msAlert.ID,
			ReceiverType: string(profile.ReceiverEmail),
			ErrMsg:       "first error",
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Equal(1, resp.Updated)

	// nlog1 (最早) 应被更新
	updated1 := s.Client.Nlog.GetX(ctx, nlog1.ID)
	s.Equal("first error", updated1.ErrMsg)

	// nlog2 (较新) 不应被影响
	updated2 := s.Client.Nlog.GetX(ctx, nlog2.ID)
	s.Empty(updated2.ErrMsg)

	// 再次请求, 应更新 nlog2 (下一个未处理的)
	gc1b, _ := gin.CreateTestContext(httptest.NewRecorder())
	gc1b.Request = httptest.NewRequest(http.MethodPost, "/nlogs/update", nil)
	gc1b.Request = gc1b.Request.WithContext(ctx)
	resp, err = s.server.UpdateNlog(gc1b, &UpdateNlogRequest{
		NlogUpdate: NlogUpdate{
			AlertID:      msAlert.ID,
			ReceiverType: string(profile.ReceiverEmail),
			ErrMsg:       "second error",
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Equal(1, resp.Updated)

	updated2 = s.Client.Nlog.GetX(ctx, nlog2.ID)
	s.Equal("second error", updated2.ErrMsg)

	// 已处理的 Nlog (err_msg 有值) 不应被 alertID+receiverType 匹配
	// 先给 nlog1 也设置 err_msg, 模拟全部已处理
	s.Client.Nlog.UpdateOneID(nlog1.ID).SetErrMsg("already processed").ExecX(ctx)

	gcSkip, _ := gin.CreateTestContext(httptest.NewRecorder())
	gcSkip.Request = httptest.NewRequest(http.MethodPost, "/nlogs/update", nil)
	gcSkip.Request = gcSkip.Request.WithContext(ctx)
	resp, err = s.server.UpdateNlog(gcSkip, &UpdateNlogRequest{
		NlogUpdate: NlogUpdate{
			AlertID:      msAlert.ID,
			ReceiverType: string(profile.ReceiverEmail),
			ErrMsg:       "should not overwrite",
		},
	})
	s.Require().NoError(err)
	s.Nil(resp, "should return nil when all matching Nlogs already have err_msg")

	// nlog2 的 ErrMsg 不应被覆盖
	updated2After := s.Client.Nlog.GetX(ctx, nlog2.ID)
	s.Equal("second error", updated2After.ErrMsg)

	// 不存在的 alertID 应返回 nil (404)
	gc2, _ := gin.CreateTestContext(httptest.NewRecorder())
	gc2.Request = httptest.NewRequest(http.MethodPost, "/nlogs/update", nil)
	gc2.Request = gc2.Request.WithContext(ctx)
	resp, err = s.server.UpdateNlog(gc2, &UpdateNlogRequest{
		NlogUpdate: NlogUpdate{
			AlertID:      99999,
			ReceiverType: string(profile.ReceiverEmail),
			ErrMsg:       "should not match",
		},
	})
	s.Require().NoError(err)
	s.Nil(resp)

	// 通过 messageId 精确匹配
	nlogWithMsg := s.Client.Nlog.Create().
		SetTenantID(1).SetReceiver("email").SetGroupKey("test-group-msg").
		SetReceiverType(profile.ReceiverEmail).SetIdx(0).
		SetMessageID("test-trace-id-123").
		SetExpiresAt(time.Now().Add(24*time.Hour)).SetSendAt(time.Now()).
		SaveX(ctx)

	gc3, _ := gin.CreateTestContext(httptest.NewRecorder())
	gc3.Request = httptest.NewRequest(http.MethodPost, "/nlogs/update", nil)
	gc3.Request = gc3.Request.WithContext(ctx)
	resp, err = s.server.UpdateNlog(gc3, &UpdateNlogRequest{
		NlogUpdate: NlogUpdate{
			MessageId: "test-trace-id-123",
			ErrMsg:    "callback error from cloud provider",
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Equal(1, resp.Updated)

	updated3 := s.Client.Nlog.GetX(ctx, nlogWithMsg.ID)
	s.Equal("callback error from cloud provider", updated3.ErrMsg)

	// 不存在的 messageId 应返回 nil (404)
	gc4, _ := gin.CreateTestContext(httptest.NewRecorder())
	gc4.Request = httptest.NewRequest(http.MethodPost, "/nlogs/update", nil)
	gc4.Request = gc4.Request.WithContext(ctx)
	resp, err = s.server.UpdateNlog(gc4, &UpdateNlogRequest{
		NlogUpdate: NlogUpdate{
			MessageId: "non-existent-id",
			ErrMsg:    "should not match",
		},
	})
	s.Require().NoError(err)
	s.Nil(resp)
}
