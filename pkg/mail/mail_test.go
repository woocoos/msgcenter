package mail

import (
	"bytes"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func newTestEmail(t *testing.T) *Email {
	t.Helper()
	msg := NewEmailMsg().SetFrom("from@example.com").AddTo("to@example.com").SetSubject("test")
	if msg.err != nil {
		t.Fatalf("build email: %v", msg.err)
	}
	return msg
}

func assertEncodedDisposition(t *testing.T, s string) {
	t.Helper()
	// RFC 2047 encoded-word (charset "utf-8" as passed to QEncoding).
	if !strings.Contains(s, `filename="`+mime.QEncoding.Encode("utf-8", "中文附件.txt")+`"`) {
		t.Errorf("missing RFC 2047 encoded-word in:\n%s", s)
	}
}

func buildMime(t *testing.T, msg *Email) string {
	t.Helper()
	raw, err := msg.e.Bytes()
	if err != nil {
		t.Fatalf("build mime: %v", err)
	}
	return string(raw)
}

func TestAttachNonASCIIFilename(t *testing.T) {
	msg := newTestEmail(t)
	if _, err := msg.Attach(bytes.NewReader([]byte("content")), "中文附件.txt", "text/plain"); err != nil {
		t.Fatalf("attach: %v", err)
	}
	assertEncodedDisposition(t, buildMime(t, msg))
}

func TestAttachFileNonASCIIFilename(t *testing.T) {
	path := filepath.Join(t.TempDir(), "中文附件.txt")
	if err := os.WriteFile(path, []byte("content"), 0o644); err != nil {
		t.Fatalf("write temp file: %v", err)
	}
	msg := newTestEmail(t)
	if _, err := msg.AttachFile(path); err != nil {
		t.Fatalf("attach file: %v", err)
	}
	assertEncodedDisposition(t, buildMime(t, msg))
}

func TestAttachASCIIFilename(t *testing.T) {
	msg := newTestEmail(t)
	if _, err := msg.Attach(bytes.NewReader([]byte("content")), "report.txt", "text/plain"); err != nil {
		t.Fatalf("attach: %v", err)
	}
	if !strings.Contains(buildMime(t, msg), `filename="report.txt"`) {
		t.Errorf("ASCII filename changed:\n%s", buildMime(t, msg))
	}
}
