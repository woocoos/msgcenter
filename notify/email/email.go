package email

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/mail"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
	"strconv"
	"strings"
)

// Notifier email notifier
//
// tmpl include all of receiver's template.
type Notifier struct {
	config        *profile.EmailConfig
	tmpl          *template.Template
	hostname      string
	customTplFunc notify.CustomerConfigFunc[profile.EmailConfig]
}

func (n *Notifier) SendResolved() bool {
	return n.config.SendResolved
}

func New(cfg *profile.EmailConfig, tmpl *template.Template, fn notify.CustomerConfigFunc[profile.EmailConfig],
) (*Notifier, error) {
	return &Notifier{
		config:        cfg,
		tmpl:          tmpl,
		customTplFunc: fn,
	}, nil
}

func (n *Notifier) getPassword() (string, error) {
	return n.config.AuthPassword, nil
}

// CustomConfig returns a custom config for the notifier.
func (n *Notifier) CustomConfig(ctx context.Context) (*profile.EmailConfig, error) {
	if n.customTplFunc == nil {
		return n.config, nil
	}
	labels, ok := notify.GroupLabels(ctx)
	if !ok {
		return n.config, nil
	}
	cfg := n.config.Clone()
	err := n.customTplFunc(ctx, cfg, labels)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Notify implements the Notifier interface.
//
// It should load customer config from DB and render the template every called.
// See service.overrideEmailConfig for more details
func (n *Notifier) Notify(ctx context.Context, alerts ...*alert.Alert) (retry bool, err error) {
	email := mail.NewEmailMsg()
	data := notify.GetTemplateData(ctx, n.tmpl, alerts)
	tmpl := notify.TmplText(n.tmpl, data, &err)
	// use custom template setting to render the email
	config, err := n.CustomConfig(ctx)
	if err != nil {
		return false, err
	}
	from := tmpl(config.From)
	if err != nil {
		return false, fmt.Errorf("execute 'from' template: %w", err)
	}
	email.SetFrom(from)

	to := tmpl(config.To)
	if err != nil {
		return false, fmt.Errorf("execute 'to' template: %w", err)
	}
	email.AddTo(strings.Split(to, ",")...)

	sub := tmpl(config.Subject)
	if err != nil {
		return false, fmt.Errorf("execute 'subject' template: %w", err)
	}
	email.SetSubject(sub)

	// choose text format as default
	if len(config.Text) > 0 {
		body, err := n.tmpl.ExecuteTextString(config.Text, data)
		if err != nil {
			return false, fmt.Errorf("execute text template: %w", err)
		}
		email.SetText(body)
	} else if len(config.HTML) > 0 {
		body, err := n.tmpl.ExecuteHTMLString(config.HTML, data)
		if err != nil {
			return false, fmt.Errorf("execute html template: %w", err)
		}
		email.SetHTML(body)
	}

	for header, t := range config.Headers {
		switch strings.ToLower(header) {
		case "attachments":
			for _, a := range strings.Split(t, ",") {
				if _, err = email.AttachFile(a); err != nil {
					return false, err
				}
			}
		case "cc":
			email.AddCc(strings.Split(t, ",")...)
		case "bcc":
			email.AddBcc(strings.Split(t, ",")...)
		default:
			value, err := n.tmpl.ExecuteTextString(t, data)
			if err != nil {
				return false, fmt.Errorf("execute %q header template: %w", header, err)
			}
			email.SetHeader(header, value)
		}
	}

	// connection level use original config
	var (
		tlsConfig *tls.Config
		ect       mail.SMTPEncryptionType
	)
	if n.config.RequireTLS {
		// new a tls.config
		tlsConfig, err = n.config.TLSConfig.BuildTlsConfig()
		if err != nil {
			return false, fmt.Errorf("parse TLS config: %w", err)
		}
		if tlsConfig.ServerName == "" {
			tlsConfig.ServerName = n.config.SmartHost.Host
		}
		ect = mail.SMTPEncryptionTypeSTARTTLS
	}
	port, _ := strconv.Atoi(n.config.SmartHost.Port)
	pwd, err := n.getPassword()
	if err != nil {
		return false, fmt.Errorf("get password: %w", err)
	}

	client := mail.NewSMTPClient(n.config.SmartHost.Host, port)
	client.SetAuthType(mail.SMTPAuthType(n.config.AuthType)).
		SetAuthCredentials(n.config.AuthIdentity, n.config.AuthUsername, pwd).
		SetEncryptionType(ect)

	if err := client.SendMail(email, tlsConfig); err != nil {
		return false, err
	}
	return true, nil
}
