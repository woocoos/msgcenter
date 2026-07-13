package email

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
