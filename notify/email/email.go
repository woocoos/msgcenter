package email

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"net/mail"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	pkgmail "github.com/woocoos/msgcenter/pkg/mail"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
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
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost.localdomain"
	}
	return &Notifier{
		config:        cfg,
		tmpl:          tmpl,
		hostname:      hostname,
		customTplFunc: fn,
	}, nil
}

// splitEmailAddresses splits a comma-separated list of email addresses,
func splitEmailAddresses(s string) ([]string, error) {
	var addresses []string
	addrs, err := mail.ParseAddressList(s)
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		addresses = append(addresses, addr.String())
	}
	return addresses, nil
}

func (n *Notifier) getPassword() (string, error) {
	if n.config.AuthPasswordFile != "" {
		content, err := os.ReadFile(n.config.AuthPasswordFile)
		if err != nil {
			return "", fmt.Errorf("reading auth password file %q: %w", n.config.AuthPasswordFile, err)
		}
		return strings.TrimSpace(string(content)), nil
	}
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
func (n *Notifier) Notify(ctx context.Context, alerts ...*alert.Alert) (retry bool, err error) {
	email := pkgmail.NewEmailMsg()
	data := notify.GetTemplateData(ctx, n.tmpl, alerts)
	tmpl := notify.TmplText(n.tmpl, data, &err)

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
	tos, err := splitEmailAddresses(to)
	if err != nil {
		return false, fmt.Errorf("parse 'to' string: %w", err)
	}
	email.AddTo(tos...)

	sub := tmpl(config.Subject)
	if err != nil {
		return false, fmt.Errorf("execute 'subject' template: %w", err)
	}
	email.SetSubject(sub)

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
			value, err := n.tmpl.ExecuteTextString(t, data)
			if err != nil {
				return false, fmt.Errorf("execute %q header template: %w", header, err)
			}
			if value != "" {
				values, err := splitEmailAddresses(value)
				if err != nil {
					return false, fmt.Errorf("parse 'cc' string: %w", err)
				}
				email.AddCc(values...)
			}
		case "bcc":
			value, err := n.tmpl.ExecuteTextString(t, data)
			if err != nil {
				return false, fmt.Errorf("execute %q header template: %w", header, err)
			}
			if value != "" {
				values, err := splitEmailAddresses(value)
				if err != nil {
					return false, fmt.Errorf("parse 'bcc' string: %w", err)
				}
				email.AddBcc(values...)
			}
		default:
			value, err := n.tmpl.ExecuteTextString(t, data)
			if err != nil {
				return false, fmt.Errorf("execute %q header template: %w", header, err)
			}
			email.SetHeader(header, value)
		}
	}

	// Generate Message-Id if not set by user.
	if _, ok := config.Headers["Message-Id"]; !ok {
		var rnd [8]byte
		rand.Read(rnd[:])
		msgID := fmt.Sprintf("<%d.%x@%s>", time.Now().UnixNano(), rnd, n.hostname)
		email.SetHeader("Message-Id", msgID)
	}

	// Email threading: add References and In-Reply-To headers.
	if config.Threading.Enabled {
		key, keyErr := notify.ExtractGroupKey(ctx)
		if keyErr == nil {
			h := sha256.Sum256([]byte(key))
			keyHash := fmt.Sprintf("%x", h[:8])
			threadBy := ""
			if config.Threading.ThreadByDate == "daily" {
				threadBy = time.Now().Format("2006-01-02")
			}
			threadRootID := fmt.Sprintf("<alert-%s-%s@msgcenter>", keyHash, threadBy)
			email.SetHeader("References", threadRootID)
			email.SetHeader("In-Reply-To", threadRootID)
		}
	}

	// Determine TLS mode.
	useImplicitTLS := false
	if config.ForceImplicitTLS != nil {
		useImplicitTLS = *config.ForceImplicitTLS
	} else {
		port, _ := strconv.Atoi(config.SmartHost.Port)
		useImplicitTLS = port == 465
	}

	var (
		tlsConfig *tls.Config
		ect       pkgmail.SMTPEncryptionType
	)
	if useImplicitTLS {
		ect = pkgmail.SMTPEncryptionTypeSSLTLS
		tlsConfig, err = config.TLSConfig.BuildTlsConfig()
		if err != nil {
			return false, fmt.Errorf("parse TLS config: %w", err)
		}
		if tlsConfig.ServerName == "" {
			tlsConfig.ServerName = config.SmartHost.Host
		}
	} else if config.RequireTLS {
		ect = pkgmail.SMTPEncryptionTypeSTARTTLS
		tlsConfig, err = config.TLSConfig.BuildTlsConfig()
		if err != nil {
			return false, fmt.Errorf("parse TLS config: %w", err)
		}
		if tlsConfig.ServerName == "" {
			tlsConfig.ServerName = config.SmartHost.Host
		}
	}

	port, _ := strconv.Atoi(config.SmartHost.Port)
	pwd, err := n.getPassword()
	if err != nil {
		return false, fmt.Errorf("get password: %w", err)
	}

	client := pkgmail.NewSMTPClient(config.SmartHost.Host, port)
	client.SetAuthType(pkgmail.SMTPAuthType(config.AuthType)).
		SetAuthCredentials(config.AuthIdentity, config.AuthUsername, pwd).
		SetEncryptionType(ect)

	if err := client.SendMail(email, tlsConfig); err != nil {
		return false, err
	}
	return true, nil
}
