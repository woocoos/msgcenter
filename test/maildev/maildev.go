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

// deleteAllEmails deletes all emails.
func (m *MailDev) deleteAllEmails() error {
	_, _, err := m.doEmailRequest(http.MethodDelete, "/email/all")
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
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}
	return res.StatusCode, b, nil
}
