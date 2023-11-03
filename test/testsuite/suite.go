package testsuite

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service"
	"github.com/woocoos/msgcenter/test"
	"os"
	"path/filepath"
	"time"
)

var (
	alterPassWordEventName = "AlterPassword"
	SubEventName           = "SubEvent"
	WebhookEventName       = "WebhookEvent"
)

type BaseSuite struct {
	suite.Suite
	App             *woocoo.App
	Cnf             *conf.AppConfiguration
	DSN, DriverName string
	Client          *ent.Client
	AlertManager    *service.AlertManager
	redis           *miniredis.Miniredis
}

func (o *BaseSuite) Setup() error {
	o.App = initTestApp()
	o.Cnf = o.App.AppConfiguration()
	o.redis = initMiniRedis(o.Cnf)

	if o.DSN == "" && o.DriverName == "" {
		o.DriverName = "sqlite3"
		o.DSN = "file:msgcenter?mode=memory&cache=shared&_fk=1"
	}
	client, err := open(context.Background(), o.DriverName, o.DSN)
	if err != nil {
		return err
	}
	o.Client = client
	initDatabase(context.Background(), o.Client)

	// alert
	metrics.BuildGlobal()

	o.AlertManager, err = service.NewAlertManager(o.App, service.WithClient(o.Client))
	o.Require().NoError(err)
	return nil
}

func (o *BaseSuite) NewTestCtx() context.Context {
	ctx := ent.NewContext(context.Background(), o.Client)
	// with identity
	ctx = security.WithContext(ctx, security.NewGenericPrincipalByClaims(jwt.MapClaims{"sub": "1"}))
	ctx = identity.WithTenantID(ctx, 1)
	return ctx
}

func initTestApp() *woocoo.App {
	file := filepath.Join(test.BaseDir(), "testdata", "etc", "app.yaml")
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	app := woocoo.New(woocoo.WithAppConfiguration(
		conf.NewFromBytes(bs, conf.WithBaseDir(test.BaseDir()))),
	)
	cnf := app.AppConfiguration()
	cnf.Parser().Set("alertManager.storage.path", filepath.Join(test.BaseDir(), "testdata", "tmp"))
	cnf.Parser().Set("alertManager.config.file", filepath.Join(test.BaseDir(), "testdata", "etc", "alertmanager.yaml"))
	return app
}

func open(ctx context.Context, driverName, dsn string) (*ent.Client, error) {
	client, err := ent.Open(driverName, dsn,
		ent.Debug(), ent.AlternateSchema(ent.SchemaConfig{
			User:        "portal",
			OrgRoleUser: "portal",
		}),
	)
	if err != nil {
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		return nil, err
	}
	return client, nil
}

func initMiniRedis(cnf *conf.AppConfiguration) *miniredis.Miniredis {
	db, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	cnf.Parser().Set("store.redis.addrs", []string{db.Addr()})
	return db
}

func initDatabase(ctx context.Context, client *ent.Client) {
	ctx = identity.WithTenantID(ctx, 1)
	client.MsgType.Create().SetName("alert").SetID(1).SetStatus(typex.SimpleStatusActive).SetCreatedBy(1).
		SetAppID(1).SetCategory("账户安全").SetCanSubs(true).SetCanCustom(true).SaveX(ctx)
	client.MsgEvent.Create().SetID(1).SetMsgTypeID(1).SetName("AlterPassword").SetStatus(typex.SimpleStatusActive).
		SetCreatedBy(1).SetModes("email,internal").
		SetRoute(&profile.Route{
			Name:     alterPassWordEventName,
			Receiver: "email",
			Matchers: label.Matchers{
				{Type: label.MatchEqual, Name: "app", Value: "1"},
				{Name: label.AlertNameLabel, Value: alterPassWordEventName, Type: label.MatchEqual},
			},
			Routes: []*profile.Route{
				{
					Matchers: label.Matchers{
						{Name: "receiver", Value: "email", Type: label.MatchRegexp},
					},
					Receiver: "email",
					Continue: true,
				},
				{
					Matchers: label.Matchers{
						{Name: "receiver", Value: "internal", Type: label.MatchRegexp},
					},
					Receiver: "internal",
				},
			},
		}).SaveX(ctx)
	client.MsgTemplate.Create().SetMsgTypeID(1).SetEventID(1).SetTenantID(1).SetName(alterPassWordEventName).SetCreatedBy(1).
		SetStatus(typex.SimpleStatusActive).SetFormat(msgtemplate.FormatTxt).SetReceiverType(profile.ReceiverEmail).SetTo(`{{ template "email.to" . }}`).
		SetSubject(`{{ with .CommonAnnotations }}{{.uid}}{{end}}密码到期提醒`).SetCc(`{{ template "email.cc" . }}`).
		SetBcc(`{{ template "email.bcc" . }}`).SetFrom(`custom <test@localhost>`).
		SetBody(`{{ template "1.alterpwd.txt" . }}`).SetAttachments([]string{"1/alterpwd.tmpl"}).
		SetTpl("1/alterpwd.tmpl").SaveX(ctx)

	client.MsgChannel.Create().SetName("email").SetStatus(typex.SimpleStatusActive).SetCreatedBy(1).
		SetTenantID(1).SetReceiverType(profile.ReceiverEmail).
		SetReceiver(&profile.Receiver{
			Name: "email",
			EmailConfigs: []*profile.EmailConfig{
				{
					SmartHost: profile.HostPort{Host: "localhost", Port: "1025"},
					To:        `{{ template "email.to" . }}`,
					From:      "serviceSuite@localhost",
				},
			},
		}).SaveX(ctx)
	client.MsgTemplate.Create().SetMsgTypeID(1).SetEventID(1).SetTenantID(1).SetName(alterPassWordEventName).SetCreatedBy(1).
		SetStatus(typex.SimpleStatusActive).SetFormat(msgtemplate.FormatTxt).SetReceiverType(profile.ReceiverEmail).SetTo(`{{ template "email.to" . }}`).
		SetSubject(`{{ with .CommonAnnotations }}{{.uid}}{{end}}密码到期提醒`).SetCc(`{{ template "email.cc" . }}`).
		SetBcc(`{{ template "email.bcc" . }}`).SetFrom(`custom <test@localhost>`).
		SetBody(`{{ template "1.alterpwd.txt" . }}`).SetAttachments([]string{"1/alterpwd.tmpl"}).
		SetTpl("1/alterpwd.tmpl").SaveX(ctx)

	client.MsgType.Create().SetName("alert1").SetID(2).SetStatus(typex.SimpleStatusActive).SetCreatedBy(1).
		SetAppID(1).SetCategory("订阅类型").SetCanSubs(true).SetCanCustom(true).SaveX(ctx)

	client.MsgEvent.Create().SetID(2).SetMsgTypeID(2).SetName(SubEventName).SetStatus(typex.SimpleStatusActive).
		SetCreatedBy(1).SetModes("email,internal").
		SetRoute(&profile.Route{
			Name:     SubEventName,
			Receiver: "email",
			Matchers: label.Matchers{
				{Type: label.MatchEqual, Name: "app", Value: "1"},
				{Name: "alertname", Value: SubEventName, Type: label.MatchEqual},
			},
			Routes: []*profile.Route{
				{
					Matchers: label.Matchers{
						{Name: "receiver", Value: "email", Type: label.MatchRegexp},
					},
					Receiver: "email",
					Continue: true,
				},
				{
					Matchers: label.Matchers{
						{Name: "receiver", Value: "internal", Type: label.MatchRegexp},
					},
					Receiver: "internal",
				},
			},
		}).SaveX(ctx)
	client.MsgTemplate.Create().SetMsgTypeID(2).SetEventID(2).SetTenantID(1).SetName(SubEventName).SetCreatedBy(1).
		SetStatus(typex.SimpleStatusActive).SetFormat(msgtemplate.FormatTxt).SetReceiverType(profile.ReceiverEmail).SetTo(`{{ template "email.to" . }}`).
		SetSubject(`订阅事件提醒`).SetCc(`{{ template "email.cc" . }}`).
		SetBcc(`{{ template "email.bcc" . }}`).SetFrom(`custom <test@localhost>`).
		SetBody(`{{ template "1.subevent.txt" . }}`).
		SetTpl("1/subevent.tmpl").SaveX(ctx)

	client.MsgChannel.Create().SetName("webhook").SetStatus(typex.SimpleStatusActive).SetCreatedBy(1).
		SetTenantID(1).SetReceiverType(profile.ReceiverWebhook).
		SetReceiver(&profile.Receiver{
			Name: "webhook",
			WebhookConfigs: []*profile.WebhookConfig{
				{
					URL: &profile.URL{Host: "localhost:5001", Scheme: "http", Path: "/webhook"},
				},
			},
		}).SaveX(ctx)
	client.MsgEvent.Create().SetID(3).SetMsgTypeID(1).SetName(WebhookEventName).SetStatus(typex.SimpleStatusActive).
		SetCreatedBy(1).SetModes("webhook").SaveX(ctx)
	client.MsgTemplate.Create().SetMsgTypeID(1).SetEventID(3).SetTenantID(1).SetName("WebhookTemplate").SetCreatedBy(1).
		SetStatus(typex.SimpleStatusActive).SetFormat(msgtemplate.FormatTxt).SetReceiverType(profile.ReceiverWebhook).
		SetSubject(`dingtalk`).
		SetBody(`{{ template "dingtalk.content" . }}`).
		SaveX(ctx)

	client.User.Create().SetID(1).SetDisplayName("admin").SetEmail("admin@localhost").
		SetPrincipalName("admin").SetMobile("13800138000").SaveX(ctx)
	client.User.Create().SetID(2).SetDisplayName("user").SetEmail("user@localhost").
		SetPrincipalName("user").SetMobile("13800138001").SaveX(ctx)
	client.User.Create().SetID(3).SetDisplayName("nobody").SetEmail("nobody@localhost").
		SetPrincipalName("nobody").SetMobile("13800138002").SaveX(ctx)
	client.OrgRoleUser.Create().SetID(1).SetOrgID(1).SetUserID(1).SetOrgRoleID(12).SetOrgUserID(3).
		SaveX(ctx)
	client.OrgRoleUser.Create().SetID(2).SetOrgID(1).SetUserID(2).SetOrgRoleID(13).SetOrgUserID(4).
		SaveX(ctx)
	client.MsgSubscriber.Create().SetID(1).SetMsgTypeID(2).SetTenantID(1).SetUserID(1).SetCreatedBy(1).
		SaveX(ctx)
	client.MsgSubscriber.Create().SetID(2).SetMsgTypeID(2).SetTenantID(1).SetOrgRoleID(13).SetCreatedBy(1).
		SaveX(ctx)
	client.MsgSubscriber.Create().SetID(3).SetMsgTypeID(2).SetTenantID(1).SetUserID(3).SetExclude(true).
		SetCreatedBy(1).SaveX(ctx)

	client.MsgInternal.Create().SetID(1).SetTenantID(1).SetSubject("subject1").SetBody("body1").SetFormat("text").
		SetCategory("订阅类型").SetCreatedBy(1).SaveX(ctx)
	client.MsgInternalTo.Create().SetID(1).SetTenantID(1).SetCreatedAt(time.Now()).SetMsgInternalID(1).
		SetUserID(1).SaveX(ctx)
	client.MsgInternal.Create().SetID(2).SetTenantID(1).SetSubject("subject2").SetBody("body2").SetFormat("text").
		SetCategory("订阅类型").SetCreatedBy(1).SaveX(ctx)
	client.MsgInternalTo.Create().SetID(2).SetTenantID(1).SetCreatedAt(time.Now()).SetMsgInternalID(2).
		SetUserID(1).SaveX(ctx)
}
