package email

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/mail"
	neturl "net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	pkgmail "github.com/woocoos/msgcenter/pkg/mail"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
)

const (
	// maxAttachmentSize limits the maximum size of a single attachment downloaded from URL.
	maxAttachmentSize = 50 << 20 // 50 MB
	// attachmentDownloadTimeout is the timeout for downloading a single attachment from URL.
	attachmentDownloadTimeout = 30 * time.Second
	// dynamicAttachmentAnnotation is the annotation key for dynamic attachment paths.
	// Value is a semicolon-separated list of file paths or HTTP(S) URLs.
	// Semicolon is used instead of comma to avoid conflicts with commas in URLs or file paths.
	dynamicAttachmentAnnotation = "__attachments__"
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
			if err := n.attachFiles(email, strings.Split(t, ",")); err != nil {
				return false, err
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

	// Attach dynamic attachments from alert annotations.
	if dynPaths := dynamicAttachmentPaths(alerts); len(dynPaths) > 0 {
		if err := n.attachFiles(email, dynPaths); err != nil {
			return false, fmt.Errorf("attach dynamic attachments: %w", err)
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

// attachFiles attaches files to the email, supporting both local file paths and HTTP(S) URLs.
func (n *Notifier) attachFiles(email *pkgmail.Email, paths []string) error {
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if strings.HasPrefix(p, "http://") || strings.HasPrefix(p, "https://") {
			if err := n.attachFromURL(email, p); err != nil {
				return fmt.Errorf("attach from URL %q: %w", p, err)
			}
		} else {
			if _, err := email.AttachFile(p); err != nil {
				return fmt.Errorf("attach file %q: %w", p, err)
			}
		}
	}
	return nil
}

// attachFromURL downloads a file from the given URL and attaches it to the email.
func (n *Notifier) attachFromURL(email *pkgmail.Email, rawURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), attachmentDownloadTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	reader := io.LimitReader(resp.Body, maxAttachmentSize)
	filename := filenameFromResponse(resp, rawURL)
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err = email.Attach(reader, filename, contentType)
	return err
}

// filenameFromResponse extracts the attachment filename from the HTTP response.
// It prefers Content-Disposition header, falling back to the URL path base.
func filenameFromResponse(resp *http.Response, rawURL string) string {
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		if _, params, err := mime.ParseMediaType(cd); err == nil {
			if name, ok := params["filename"]; ok && name != "" {
				return name
			}
		}
	}
	if u, err := neturl.Parse(rawURL); err == nil && u.Path != "" {
		if base := path.Base(u.Path); base != "." && base != "/" {
			return base
		}
	}
	return "attachment"
}

// dynamicAttachmentPaths extracts attachment paths from alert annotations.
// Multiple alerts may carry different attachments; duplicates are preserved
// and the caller is responsible for deduplication if needed.
func dynamicAttachmentPaths(alerts []*alert.Alert) []string {
	var paths []string
	seen := make(map[string]struct{})
	for _, a := range alerts {
		v, ok := a.Annotations[dynamicAttachmentAnnotation]
		if !ok || v == "" {
			continue
		}
		for _, p := range strings.Split(v, ";") {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			if _, dup := seen[p]; dup {
				continue
			}
			seen[p] = struct{}{}
			paths = append(paths, p)
		}
	}
	return paths
}
