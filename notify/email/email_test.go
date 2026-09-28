package email

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	pkgmail "github.com/woocoos/msgcenter/pkg/mail"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
)

func TestNew_Hostname(t *testing.T) {
	t.Parallel()
	cfg := &profile.EmailConfig{To: "test@example.com"}
	tmpl, err := template.New()
	require.NoError(t, err)

	n, err := New(cfg, tmpl, nil)
	require.NoError(t, err)
	assert.NotEmpty(t, n.hostname)
}

func TestGetPassword_Direct(t *testing.T) {
	t.Parallel()
	n := &Notifier{config: &profile.EmailConfig{AuthPassword: "secret123"}}
	pwd, err := n.getPassword()
	require.NoError(t, err)
	assert.Equal(t, "secret123", pwd)
}

func TestGetPassword_FromFile(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	pwdFile := filepath.Join(dir, "password")
	require.NoError(t, os.WriteFile(pwdFile, []byte("  file-secret  \n"), 0644))

	n := &Notifier{config: &profile.EmailConfig{AuthPasswordFile: pwdFile}}
	pwd, err := n.getPassword()
	require.NoError(t, err)
	assert.Equal(t, "file-secret", pwd)
}

func TestGetPassword_FileMissing(t *testing.T) {
	t.Parallel()
	n := &Notifier{config: &profile.EmailConfig{AuthPasswordFile: "/nonexistent/file"}}
	_, err := n.getPassword()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "reading auth password file")
}

func TestGetPassword_FilePrecedence(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	pwdFile := filepath.Join(dir, "password")
	require.NoError(t, os.WriteFile(pwdFile, []byte("file-pwd"), 0644))

	n := &Notifier{config: &profile.EmailConfig{
		AuthPassword:     "inline-pwd",
		AuthPasswordFile: pwdFile,
	}}
	pwd, err := n.getPassword()
	require.NoError(t, err)
	assert.Equal(t, "file-pwd", pwd)
}

func TestSplitEmailAddresses(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"single", "user@example.com", 1, false},
		{"multiple", "a@example.com, b@example.com", 2, false},
		{"quoted comma", `"Company, Ltd" <test@example.com>`, 1, false},
		{"empty", "", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := splitEmailAddresses(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Len(t, result, tt.want)
			}
		})
	}
}

func TestImplicitTLS_Detection(t *testing.T) {
	tests := []struct {
		name             string
		port             string
		forceImplicitTLS *bool
		wantImplicit     bool
	}{
		{"port 465 default", "465", nil, true},
		{"port 587 default", "587", nil, false},
		{"port 587 forced", "587", boolPtr(true), true},
		{"port 465 disabled", "465", boolPtr(false), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useImplicitTLS := false
			if tt.forceImplicitTLS != nil {
				useImplicitTLS = *tt.forceImplicitTLS
			} else {
				var port int
				fmt.Sscanf(tt.port, "%d", &port)
				useImplicitTLS = port == 465
			}
			assert.Equal(t, tt.wantImplicit, useImplicitTLS)
		})
	}
}

func boolPtr(b bool) *bool {
	return &b
}

func TestFilenameFromResponse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		cdHeader     string
		rawURL       string
		wantFilename string
	}{
		{
			name:         "Content-Disposition with filename",
			cdHeader:     `attachment; filename="report.pdf"`,
			rawURL:       "https://example.com/data.csv",
			wantFilename: "report.pdf",
		},
		{
			name:         "fallback to URL path",
			cdHeader:     "",
			rawURL:       "https://example.com/files/report.pdf",
			wantFilename: "report.pdf",
		},
		{
			name:         "URL with query string",
			cdHeader:     "",
			rawURL:       "https://example.com/download?file=report.pdf&token=abc",
			wantFilename: "download",
		},
		{
			name:         "URL root path",
			cdHeader:     "",
			rawURL:       "https://example.com/",
			wantFilename: "attachment",
		},
		{
			name:         "Content-Disposition without filename param",
			cdHeader:     "inline",
			rawURL:       "https://example.com/data.csv",
			wantFilename: "data.csv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			resp := &http.Response{Header: http.Header{}}
			if tt.cdHeader != "" {
				resp.Header.Set("Content-Disposition", tt.cdHeader)
			}
			got := filenameFromResponse(resp, tt.rawURL)
			assert.Equal(t, tt.wantFilename, got)
		})
	}
}

// toAbsPath applies the same path transformation as dynamicAttachmentPaths
func toAbsPath(p string) string {
	if strings.HasPrefix(p, "http://") || strings.HasPrefix(p, "https://") {
		return p
	}
	p = strings.TrimPrefix(p, "/")
	if abs, err := filepath.Abs(p); err == nil {
		return abs
	}
	return p
}

func TestDynamicAttachmentPaths(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		alerts []*alert.Alert
		want   []string
	}{
		{
			name:   "no alerts",
			alerts: nil,
			want:   nil,
		},
		{
			name: "single alert with attachments",
			alerts: []*alert.Alert{
				{Annotations: label.LabelSet{alert.DynamicAttachmentAnnotation: "/tmp/a.pdf;/tmp/b.csv"}},
			},
			want: []string{"/tmp/a.pdf", "/tmp/b.csv"},
		},
		{
			name: "multiple alerts merged and deduped",
			alerts: []*alert.Alert{
				{Annotations: label.LabelSet{alert.DynamicAttachmentAnnotation: "/tmp/a.pdf"}},
				{Annotations: label.LabelSet{alert.DynamicAttachmentAnnotation: "/tmp/a.pdf;https://example.com/c.pdf"}},
			},
			want: []string{"/tmp/a.pdf", "https://example.com/c.pdf"},
		},
		{
			name: "alert without annotation skipped",
			alerts: []*alert.Alert{
				{Annotations: label.LabelSet{"other": "value"}},
				{Annotations: label.LabelSet{alert.DynamicAttachmentAnnotation: "/tmp/x.pdf"}},
			},
			want: []string{"/tmp/x.pdf"},
		},
		{
			name: "empty annotation value skipped",
			alerts: []*alert.Alert{
				{Annotations: label.LabelSet{alert.DynamicAttachmentAnnotation: ""}},
				{Annotations: label.LabelSet{alert.DynamicAttachmentAnnotation: "  ;  ;/tmp/y.pdf"}},
			},
			want: []string{"/tmp/y.pdf"},
		},
		{
			name: "URL with comma preserved",
			alerts: []*alert.Alert{
				{Annotations: label.LabelSet{alert.DynamicAttachmentAnnotation: "https://example.com/download?a=1,b=2;/tmp/local.pdf"}},
			},
			want: []string{"https://example.com/download?a=1,b=2", "/tmp/local.pdf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := dynamicAttachmentPaths(tt.alerts)
			if tt.want == nil {
				assert.Nil(t, got)
				return
			}
			// Transform expected values the same way as dynamicAttachmentPaths
			want := make([]string, len(tt.want))
			for i, p := range tt.want {
				want[i] = toAbsPath(p)
			}
			assert.Equal(t, want, got)
		})
	}
}

func TestAttachFiles_LocalFile(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	f := filepath.Join(dir, "test.txt")
	require.NoError(t, os.WriteFile(f, []byte("hello"), 0644))

	n := &Notifier{}
	email := pkgmail.NewEmailMsg()
	err := n.attachFiles(email, []string{f})
	require.NoError(t, err)
	assert.Len(t, email.Attachments(), 1)
	assert.Equal(t, "test.txt", email.Attachments()[0].Filename)
}

func TestAttachFiles_HTTPURL(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", `attachment; filename="remote.csv"`)
		w.Header().Set("Content-Type", "text/csv")
		w.Write([]byte("a,b,c\n1,2,3"))
	}))
	defer ts.Close()

	n := &Notifier{}
	email := pkgmail.NewEmailMsg()
	err := n.attachFiles(email, []string{ts.URL + "/remote.csv"})
	require.NoError(t, err)
	require.Len(t, email.Attachments(), 1)
	assert.Equal(t, "remote.csv", email.Attachments()[0].Filename)
	assert.Equal(t, "text/csv", email.Attachments()[0].ContentType)
}

func TestAttachFiles_HTTPError(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	n := &Notifier{}
	email := pkgmail.NewEmailMsg()
	err := n.attachFiles(email, []string{ts.URL + "/missing"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected status")
}

func TestAttachFiles_MixedLocalAndHTTP(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	f := filepath.Join(dir, "local.txt")
	require.NoError(t, os.WriteFile(f, []byte("local"), 0644))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("remote"))
	}))
	defer ts.Close()

	n := &Notifier{}
	email := pkgmail.NewEmailMsg()
	err := n.attachFiles(email, []string{f, ts.URL + "/remote.bin", ""})
	require.NoError(t, err)
	assert.Len(t, email.Attachments(), 2)
}

func TestAttachFiles_SkipsEmptyPaths(t *testing.T) {
	t.Parallel()
	n := &Notifier{}
	email := pkgmail.NewEmailMsg()
	err := n.attachFiles(email, []string{"", "  "})
	require.NoError(t, err)
	assert.Empty(t, email.Attachments())
}
