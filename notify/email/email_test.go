package email

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
	"github.com/woocoos/msgcenter/test/maildev"
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

func notifyEmail(cfg *profile.EmailConfig, server *maildev.MailDev) (*maildev.MailDevEmail, bool, error) {
	return notifyEmailWithContext(context.Background(), cfg, server)
}

// notifyEmailWithContext sends a notification with one firing alert and retrieves the
// email from the SMTP server if the notification has been successfully delivered.
func notifyEmailWithContext(ctx context.Context, cfg *profile.EmailConfig, server *maildev.MailDev) (*maildev.MailDevEmail, bool, error) {
	if cfg.Headers == nil {
		cfg.Headers = make(map[string]string)
	}
	firingAlert := &alert.Alert{
		Labels:   label.LabelSet{},
		StartsAt: time.Now(),
		EndsAt:   time.Now().Add(time.Hour),
	}
	err := server.DeleteAllEmails()
	if err != nil {
		return nil, false, err
	}
	tmpl, err := template.New()
	if err != nil {
		return nil, false, err
	}
	template.MustParse(tmpl.ParseGlob("*"))
	tmpl.ExternalURL, _ = url.Parse("http://am")
	email, _ := New(cfg, tmpl, nil)

	retry, err := email.Notify(ctx, firingAlert)
	if err != nil {
		return nil, retry, err
	}

	e, err := server.GetLastEmail()
	if err != nil {
		return nil, retry, err
	} else if e == nil {
		return nil, retry, fmt.Errorf("email not found")
	}
	return e, retry, nil
}

type EmailSuite struct {
	suite.Suite
	noAuthMailDev *maildev.MailDev
}

func (ts *EmailSuite) SetupSuite() {
	host := os.Getenv(emailNoAuthConfigVar)
	if host == "" {
		host = emailNoAuthHost
	}
	ts.noAuthMailDev = maildev.DefaultServer()
	host = os.Getenv(emailAuthConfigVar)
	if host == "" {
		host = emailAuthHost
	}
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
				cfg.Text = ""
				cfg.HTML = `{{ template "invalid" }}`
			},
			errMsg: `execute html template:`,
		},
	} {
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

			e, err := ts.noAuthMailDev.GetLastEmail()
			ts.Require().NoError(err)
			if tc.hasEmail {
				ts.Require().NotNil(e)
			} else {
				ts.Require().Nil(e)
			}
		})
	}
}
