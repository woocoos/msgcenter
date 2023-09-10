// Package mail is a mail delivery plugin based on SMTP.
package mail

// Usage:
//  email := NewEmailMsg().SetFrom("Bytebase <from@bytebase.com>").AddTo("Customer <to@bytebase.com>").SetSubject("Test Email Subject").SetBody(`
// <!DOCTYPE html>
// <html>
// <head>
// 	<title>HTML Test</title>
// </head>
// <body>
// 	<h1>This is a mail delivery test.</h1>
// </body>
// </html>
// 	`)
// 	fmt.Printf("email: %+v\n", email)
// 	client := NewSMTPClient("smtp.gmail.com", 587)
// 	client.SetAuthType(SMTPAuthTypePlain)
// 	client.SetAuthCredentials("from@bytebase.com", "nqxxxxxxxxxxxxxx")
// 	client.SetEncryptionType(SMTPEncryptionTypeSTARTTLS)
// 	if err := client.SendMail(email); err != nil {
// 		t.Fatalf("SendMail failed: %v", err)
// 	}

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/mail"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// Email is the email to be sent.
type Email struct {
	err     error
	from    string
	subject string

	e *email.Email
}

// NewEmailMsg returns a new email message.
func NewEmailMsg() *Email {
	e := &Email{
		e: email.NewEmail(),
	}
	return e
}

// SetFrom sets the from address of the SMTP client.
// Only accept the valid RFC 5322 address, e.g. "Bytebase <support@bytebase.com>".
func (e *Email) SetFrom(from string) *Email {
	if e.err != nil {
		return e
	}
	if e.from != "" {
		e.err = errors.New("from address already set")
		return e
	}

	parsedAddr, err := mail.ParseAddress(from)
	if err != nil {
		e.err = fmt.Errorf("invalid from address: %w: %s", err, from)
		return e
	}
	e.from = parsedAddr.Address
	e.e.From = parsedAddr.String()
	return e
}

// AddTo adds the to address of the SMTP client.
// Only accept the valid RFC 5322 address, e.g. "Name <user@domain.com>".
func (e *Email) AddTo(to ...string) *Email {
	if e.err != nil {
		return e
	}
	var buf []*mail.Address
	for _, toAddress := range to {
		parsedAddr, err := mail.ParseAddress(toAddress)
		if err != nil {
			e.err = fmt.Errorf("invalid to address: %w: %s", err, toAddress)
			return e
		}
		buf = append(buf, parsedAddr)
	}
	for _, addr := range buf {
		e.e.To = append(e.e.To, addr.String())
	}
	return e
}

func (e *Email) SetHeader(k, v string) *Email {
	if e.err != nil {
		return e
	}
	e.e.Headers[k] = []string{v}
	return e
}

func (e *Email) AddCc(cc ...string) *Email {
	if e.err != nil {
		return e
	}
	var buf []*mail.Address
	for _, addr := range cc {
		parsedAddr, err := mail.ParseAddress(addr)
		if err != nil {
			e.err = fmt.Errorf("invalid to address: %w: %s", err, addr)
			return e
		}
		buf = append(buf, parsedAddr)
	}
	for _, addr := range buf {
		e.e.Cc = append(e.e.Cc, addr.String())
	}
	return e
}

func (e *Email) AddBcc(bcc ...string) *Email {
	if e.err != nil {
		return e
	}
	var buf []*mail.Address
	for _, addr := range bcc {
		parsedAddr, err := mail.ParseAddress(addr)
		if err != nil {
			e.err = fmt.Errorf("invalid to address: %w: %s", err, addr)
			return e
		}
		buf = append(buf, parsedAddr)
	}
	for _, addr := range buf {
		e.e.Bcc = append(e.e.Bcc, addr.String())
	}
	return e
}

// SetSubject sets the subject of the SMTP client.
func (e *Email) SetSubject(subject string) *Email {
	if e.err != nil {
		return e
	}
	if e.subject != "" {
		e.err = errors.New("subject already set")
		return e
	}
	e.subject = subject
	e.e.Subject = subject
	return e
}

func (e *Email) SetText(text string) *Email {
	e.e.Text = []byte(text)
	return e
}

// SetHTML sets the body of the SMTP client. It must be html formatted.
func (e *Email) SetHTML(body string) *Email {
	e.e.HTML = []byte(body)
	return e
}

// AttachFile attaches the file to the email, and returns the filename of the attachment.
// Caller can use filename as content id to reference the attachment in the email body.
func (e *Email) AttachFile(filename string) (*email.Attachment, error) {
	return e.e.AttachFile(filename)
}

// SMTPAuthType is the type of SMTP authentication.
type SMTPAuthType string

const (
	// SMTPAuthTypeNone is the NONE auth type of SMTP.
	SMTPAuthTypeNone SMTPAuthType = ""
	// SMTPAuthTypePlain is the PLAIN auth type of SMTP.
	SMTPAuthTypePlain SMTPAuthType = "PLAIN"
	// SMTPAuthTypeLogin is the LOGIN auth type of SMTP.
	SMTPAuthTypeLogin SMTPAuthType = "LOGIN"
	// SMTPAuthTypeCRAMMD5 is the CRAM-MD5 auth type of SMTP.
	SMTPAuthTypeCRAMMD5 SMTPAuthType = "CRAM-MD5"
)

// Validate implements field.EnumValues interface
func (SMTPAuthType) Validate(in string) (bool, error) {
	authType := SMTPAuthType(in)
	switch authType {
	case SMTPAuthTypeNone, SMTPAuthTypePlain, SMTPAuthTypeLogin, SMTPAuthTypeCRAMMD5:
		return true, nil
	}
	return false, fmt.Errorf("invalid SMTPAuthType: %s", in)
}

// SMTPEncryptionType is the type of SMTP encryption.
type SMTPEncryptionType uint

const (
	// SMTPEncryptionTypeNone is the NONE encrypt type of SMTP.
	SMTPEncryptionTypeNone = iota
	// SMTPEncryptionTypeSSLTLS is the SSL/TLS encrypt type of SMTP.
	SMTPEncryptionTypeSSLTLS
	// SMTPEncryptionTypeSTARTTLS is the STARTTLS encrypt type of SMTP.
	SMTPEncryptionTypeSTARTTLS
)

// SMTPClient is the client of SMTP.
type SMTPClient struct {
	host           string
	port           int
	authType       SMTPAuthType
	identity       string
	username       string
	password       string
	encryptionType SMTPEncryptionType
}

// NewSMTPClient returns a new SMTP client.
func NewSMTPClient(host string, port int) *SMTPClient {
	return &SMTPClient{
		host:           host,
		port:           port,
		authType:       SMTPAuthTypeNone,
		username:       "",
		password:       "",
		encryptionType: SMTPEncryptionTypeNone,
	}
}

// SendMail sends the email.
func (c *SMTPClient) SendMail(e *Email, t *tls.Config) error {
	if e.err != nil {
		return e.err
	}

	switch c.encryptionType {
	case SMTPEncryptionTypeNone:
		return e.e.Send(fmt.Sprintf("%s:%d", c.host, c.port), c.getAuth())
	case SMTPEncryptionTypeSSLTLS:
		if t == nil {
			t = &tls.Config{}
		}
		return e.e.SendWithTLS(fmt.Sprintf("%s:%d", c.host, c.port), c.getAuth(), t)
	case SMTPEncryptionTypeSTARTTLS:
		if t == nil {
			t = &tls.Config{InsecureSkipVerify: true}
		}
		return e.e.SendWithStartTLS(fmt.Sprintf("%s:%d", c.host, c.port), c.getAuth(), t)
	default:
		return fmt.Errorf("unknown SMTP encryption type: %d", c.encryptionType)
	}
}

// SetAuthType sets the auth type of the SMTP client.
func (c *SMTPClient) SetAuthType(authType SMTPAuthType) *SMTPClient {
	c.authType = authType
	return c
}

// SetAuthCredentials sets the auth credentials of the SMTP client.
func (c *SMTPClient) SetAuthCredentials(identity, username, password string) *SMTPClient {
	c.identity = identity
	c.username = username
	c.password = password
	return c
}

func (c *SMTPClient) getAuth() smtp.Auth {
	switch c.authType {
	case SMTPAuthTypeNone:
		return nil
	case SMTPAuthTypePlain:
		return smtp.PlainAuth(c.identity, c.username, c.password, c.host)
	case SMTPAuthTypeLogin:
		return LoginAuth(c.username, c.password)
	case SMTPAuthTypeCRAMMD5:
		return smtp.CRAMMD5Auth(c.username, c.password)
	default:
		return nil
	}
}

// SetEncryptionType sets the encryption type of the SMTP client.
func (c *SMTPClient) SetEncryptionType(encryptionType SMTPEncryptionType) {
	c.encryptionType = encryptionType
}
