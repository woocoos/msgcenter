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

// splitEmailAddresses splits a comma-separated list of email addresses,
// but ignores commas inside quoted strings or angle brackets.
// For example: `"Company, Ltd" <test@example.com>, user2@example.com`
// will be split into `["Company, Ltd" <test@example.com>, user2@example.com]`
// It also fixes addresses missing angle brackets: `"Name" email@example.com` -> `"Name" <email@example.com>`
func splitEmailAddresses(s string) []string {
	var addresses []string
	var current strings.Builder
	inQuotes := false
	inAngleBrackets := false

	for _, r := range s {
		switch r {
		case '"':
			inQuotes = !inQuotes
			current.WriteRune(r)
		case '<':
			inAngleBrackets = true
			current.WriteRune(r)
		case '>':
			inAngleBrackets = false
			current.WriteRune(r)
		case ',':
			if inQuotes || inAngleBrackets {
				// Inside quotes or angle brackets, comma is part of the address
				current.WriteRune(r)
			} else {
				// Outside quotes and angle brackets, comma is a separator
				addr := strings.TrimSpace(current.String())
				if addr != "" {
					addresses = append(addresses, fixAddressFormat(addr))
				}
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}
	}

	// Add the last address
	addr := strings.TrimSpace(current.String())
	if addr != "" {
		addresses = append(addresses, fixAddressFormat(addr))
	}

	return addresses
}

// fixAddressFormat fixes common email address format issues.
// For example: converts `"Name" email@example.com` to `"Name" <email@example.com>`.
func fixAddressFormat(addr string) string {
	addr = strings.TrimSpace(addr)
	// If the address already contains angle brackets, it's likely already in correct format
	if strings.Contains(addr, "<") && strings.Contains(addr, ">") {
		return addr
	}

	// Try to find pattern: "Name" email@domain.com (quoted name followed by email without angle brackets)
	if len(addr) > 0 && addr[0] == '"' {
		// Find the closing quote
		endQuote := strings.Index(addr[1:], `"`)
		if endQuote > 0 {
			endQuote += 1 // Adjust for the offset
			remaining := strings.TrimSpace(addr[endQuote+1:])
			if remaining != "" && strings.Contains(remaining, "@") {
				// Check if remaining looks like an email address (has @ and no spaces)
				parts := strings.Fields(remaining)
				if len(parts) == 1 && strings.Contains(parts[0], "@") {
					// Format: "Name" <email>
					return fmt.Sprintf("%s <%s>", addr[:endQuote+1], parts[0])
				}
			}
		}
	}

	return addr
}

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
	email.AddTo(splitEmailAddresses(to)...)

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
			value, err := n.tmpl.ExecuteTextString(t, data)
			if err != nil {
				return false, fmt.Errorf("execute %q header template: %w", header, err)
			}
			if value != "" {
				email.AddCc(splitEmailAddresses(value)...)
			}
		case "bcc":
			value, err := n.tmpl.ExecuteTextString(t, data)
			if err != nil {
				return false, fmt.Errorf("execute %q header template: %w", header, err)
			}
			if value != "" {
				email.AddBcc(splitEmailAddresses(value)...)
			}
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
