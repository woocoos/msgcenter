package maildev

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	Scheme = "http"
	Host   = "localhost:8025"
)

// MailDev is a client for the MailDev server.
type (
	MailDev struct {
		*url.URL
	}
	MailDevEmail struct {
		Attachments int
		Bcc         []map[string]string
		Cc          []map[string]string
		Created     time.Time
		From        map[string]string
		ID          string
		Read        bool
		Size        int
		Subject     string
		Tags        []string
		To          []map[string]string
	}
)

func DefaultServer() *MailDev {
	return &MailDev{
		URL: &url.URL{
			Scheme: Scheme,
			Host:   Host,
		},
	}
}

// GetLastEmail returns the last received email.
func (m *MailDev) GetLastEmail() (*MailDevEmail, error) {
	code, b, err := m.doEmailRequest(http.MethodGet, "/api/v1/messages")
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("expected status OK, got %d", code)
	}
	var emails struct {
		Messages []MailDevEmail `json:"messages"`
	}
	err = json.Unmarshal(b, &emails)
	if err != nil {
		return nil, err
	}
	if len(emails.Messages) == 0 {
		return nil, nil
	}
	return &emails.Messages[0], nil
}

// DeleteAllEmails deletes all emails.
func (m *MailDev) DeleteAllEmails() error {
	_, _, err := m.doEmailRequest(http.MethodDelete, "/api/v1/messages")
	return err
}

// doEmailRequest makes a request to the MailDev API.
func (m *MailDev) doEmailRequest(method, path string) (int, []byte, error) {
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
	if res.StatusCode != 200 {
		return res.StatusCode, nil, fmt.Errorf("expected status OK, got %d", res.StatusCode)
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}
	return res.StatusCode, b, nil
}
