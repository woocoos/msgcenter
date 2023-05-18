// Some tests require a running mail catcher. We use MailDev for this purpose,
// it can work without or with authentication (LOGIN only). It exposes a REST
// API which we use to retrieve and check the sent emails.
//
// Those tests are only executed when specific environment variables are set,
// otherwise they are skipped. The tests must be run by the CI.
//
// To run the tests locally, you should start 2 MailDev containers:
//
// $ docker run --rm -p 10080:1080 -p 1025:1025 -p 10090:1090 --entrypoint bin/maildev djfarrelly/maildev -v -w 1090
// $ docker run --rm -p 10081:1080 -p 1026:1025 -p 10091:1090 --entrypoint bin/maildev djfarrelly/maildev --incoming-user user --incoming-pass pass -v -w 1090
//
// $ EMAIL_NO_AUTH_CONFIG=testdata/noauth.yml EMAIL_AUTH_CONFIG=testdata/auth.yml make
//
// See also https://github.com/djfarrelly/MailDev for more details.
package email

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

const (
	emailNoAuthConfigVar = "EMAIL_NO_AUTH_CONFIG"
	emailAuthConfigVar   = "EMAIL_AUTH_CONFIG"

	emailNoAuthHost = "localhost:1080"
	emailAuthHost   = "localhost:1081"
	emailTo         = "alerts@example.com"
	emailFrom       = "alertmanager@example.com"
)

// email represents an email returned by the MailDev REST API.
// See https://github.com/djfarrelly/MailDev/blob/master/docs/rest.md.
type email struct {
	To      []map[string]string
	From    []map[string]string
	Subject string
	HTML    *string
	Text    *string
	Headers map[string]string
}

// mailDev is a client for the MailDev server.
type mailDev struct {
	*url.URL
}

// getLastEmail returns the last received email.
func (m *mailDev) getLastEmail() (*email, error) {
	code, b, err := m.doEmailRequest(http.MethodGet, "/email")
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("expected status OK, got %d", code)
	}

	var emails []email
	err = json.Unmarshal(b, &emails)
	if err != nil {
		return nil, err
	}
	if len(emails) == 0 {
		return nil, nil
	}
	return &emails[len(emails)-1], nil
}

// deleteAllEmails deletes all emails.
func (m *mailDev) deleteAllEmails() error {
	_, _, err := m.doEmailRequest(http.MethodDelete, "/email/all")
	return err
}

// doEmailRequest makes a request to the MailDev API.
func (m *mailDev) doEmailRequest(method, path string) (int, []byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s://%s%s", m.Scheme, m.Host, path), nil)
	if err != nil {
		return 0, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	req = req.WithContext(ctx)
	defer cancel()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}
	return res.StatusCode, b, nil
}

func notifyEmail(cfg *profile.EmailConfig, server *mailDev) (*email, bool, error) {
	return notifyEmailWithContext(context.Background(), cfg, server)
}

// notifyEmailWithContext sends a notification with one firing alert and retrieves the
// email from the SMTP server if the notification has been successfully delivered.
func notifyEmailWithContext(ctx context.Context, cfg *profile.EmailConfig, server *mailDev) (*email, bool, error) {
	if cfg.Headers == nil {
		cfg.Headers = make(map[string]string)
	}
	firingAlert := &alert.Alert{
		Labels:   label.LabelSet{},
		StartsAt: time.Now(),
		EndsAt:   time.Now().Add(time.Hour),
	}
	err := server.deleteAllEmails()
	if err != nil {
		return nil, false, err
	}
	tmpl, err := template.New()
	if err != nil {
		return nil, false, err
	}
	template.MustParse(tmpl.ParseGlob("*"))
	tmpl.ExternalURL, _ = url.Parse("http://am")
	email := NewEmail(cfg, tmpl, nil)

	retry, err := email.Notify(ctx, firingAlert)
	if err != nil {
		return nil, retry, err
	}

	e, err := server.getLastEmail()
	if err != nil {
		return nil, retry, err
	} else if e == nil {
		return nil, retry, fmt.Errorf("email not found")
	}
	return e, retry, nil
}

type EmailSuite struct {
	suite.Suite
	authMailDev   *mailDev
	noAuthMailDev *mailDev
}

func (ts *EmailSuite) SetupSuite() {
	host := os.Getenv(emailNoAuthConfigVar)
	if host == "" {
		host = emailNoAuthHost
	}
	ts.noAuthMailDev = &mailDev{URL: &url.URL{Scheme: "http", Host: host}}
	host = os.Getenv(emailAuthConfigVar)
	if host == "" {
		host = emailAuthHost
	}
	ts.authMailDev = &mailDev{URL: &url.URL{Scheme: "http", Host: host}}
}

func TestEmailSuite(t *testing.T) {
	suite.Run(t, new(EmailSuite))
}

func (ts *EmailSuite) TestEmailNotifyWithErrors() {
	for _, tc := range []struct {
		title     string
		updateCfg func(*profile.EmailConfig)

		errMsg   string
		hasEmail bool
	}{
		{
			title: "invalid 'from' template",
			updateCfg: func(cfg *profile.EmailConfig) {
				cfg.From = `{{ template "invalid" }}`
			},
			errMsg: "execute 'from' template:",
		},
		{
			title: "invalid 'from' address",
			updateCfg: func(cfg *profile.EmailConfig) {
				cfg.From = `xxx`
			},
			errMsg: "invalid from address:",
		},
		{
			title: "invalid 'to' template",
			updateCfg: func(cfg *profile.EmailConfig) {
				cfg.To = `{{ template "invalid" }}`
			},
			errMsg: "execute 'to' template:",
		},
		{
			title: "invalid 'to' address",
			updateCfg: func(cfg *profile.EmailConfig) {
				cfg.To = `xxx`
			},
			errMsg: "invalid to address:",
		},
		{
			title: "invalid 'subject' template",
			updateCfg: func(cfg *profile.EmailConfig) {
				cfg.Headers["Subject"] = `{{ template "invalid" }}`
			},
			errMsg: `execute "Subject" header template:`,
		},
		{
			title: "invalid 'text' template",
			updateCfg: func(cfg *profile.EmailConfig) {
				cfg.Text = `{{ template "invalid" }}`
			},
			errMsg: `execute text template:`,
		},
		{
			title: "invalid 'html' template",
			updateCfg: func(cfg *profile.EmailConfig) {
				cfg.HTML = `{{ template "invalid" }}`
			},
			errMsg: `execute html template:`,
		},
	} {
		tc := tc
		ts.Run(tc.title, func() {
			if len(tc.errMsg) == 0 {
				ts.T().Fatal("please define the expected error message")
				return
			}

			emailCfg := &profile.EmailConfig{
				SmartHost: profile.HostPort{Port: "1025", Host: "localhost"},
				To:        emailTo,
				From:      emailFrom,
				HTML:      "HTML body",
				Text:      "Text body",
				Headers: map[string]string{
					"Subject": "{{ len .Alerts }} {{ .Status }} alert(s)",
				},
			}
			if tc.updateCfg != nil {
				tc.updateCfg(emailCfg)
			}

			_, retry, err := notifyEmail(emailCfg, ts.noAuthMailDev)
			ts.Require().Error(err)
			ts.Require().Contains(err.Error(), tc.errMsg)
			ts.Require().Equal(false, retry)

			e, err := ts.noAuthMailDev.getLastEmail()
			ts.Require().NoError(err)
			if tc.hasEmail {
				ts.Require().NotNil(e)
			} else {
				ts.Require().Nil(e)
			}
		})
	}
}
