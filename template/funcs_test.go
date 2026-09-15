package template

import (
	"testing"
	"text/template"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func execTemplate(t *testing.T, tmplText string, data any) string {
	t.Helper()
	tmpl, err := template.New("test").Funcs(DefaultFuncs).Parse(tmplText)
	require.NoError(t, err)
	var buf []byte
	w := &writer{buf: &buf}
	require.NoError(t, tmpl.Execute(w, data))
	return string(buf)
}

type writer struct {
	buf *[]byte
}

func (w *writer) Write(p []byte) (n int, err error) {
	*w.buf = append(*w.buf, p...)
	return len(p), nil
}

func TestDefaultFuncs_Registered(t *testing.T) {
	t.Parallel()
	expected := []string{
		"toUpper", "toLower", "title", "trimSpace", "join", "match",
		"safeHtml", "safeUrl", "urlUnescape", "reReplaceAll", "stringSlice",
		"markdown", "toJSON", "now", "since", "date", "tz",
		"toDate", "mustToDate", "humanizeDuration", "list", "append", "dict",
	}
	for _, name := range expected {
		_, ok := DefaultFuncs[name]
		assert.True(t, ok, "function %q not registered", name)
	}
}

func TestDate(t *testing.T) {
	t.Parallel()
	ts := time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC)
	result := execTemplate(t, `{{ date "2006-01-02 15:04" . }}`, ts)
	assert.Equal(t, "2024-06-15 10:30", result)
}

func TestTz(t *testing.T) {
	t.Parallel()
	ts := time.Date(2024, 6, 15, 10, 0, 0, 0, time.UTC)
	result := execTemplate(t, `{{ tz "Asia/Shanghai" . | date "15:04" }}`, ts)
	assert.Equal(t, "18:00", result)
}

func TestTz_InvalidLocation(t *testing.T) {
	t.Parallel()
	ts := time.Date(2024, 6, 15, 10, 0, 0, 0, time.UTC)
	tmpl, err := template.New("test").Funcs(DefaultFuncs).Parse(`{{ tz "Invalid/Zone" . }}`)
	require.NoError(t, err)
	err = tmpl.Execute(&writer{}, ts)
	assert.Error(t, err)
}

func TestNow(t *testing.T) {
	t.Parallel()
	before := time.Now()
	result := execTemplate(t, `{{ now | date "2006" }}`, nil)
	after := time.Now()
	year := result
	assert.GreaterOrEqual(t, year, before.Format("2006"))
	assert.LessOrEqual(t, year, after.Format("2006"))
}

func TestSince(t *testing.T) {
	t.Parallel()
	past := time.Now().Add(-2 * time.Hour)
	result := execTemplate(t, `{{ since . | humanizeDuration }}`, past)
	assert.Contains(t, result, "h")
}

func TestHumanizeDuration(t *testing.T) {
	tests := []struct {
		name     string
		template string
		data     any
		expected string
	}{
		{"zero", `{{ humanizeDuration 0 }}`, nil, "0s"},
		{"seconds", `{{ humanizeDuration 45 }}`, nil, "45s"},
		{"minutes", `{{ humanizeDuration 125 }}`, nil, "2m 5s"},
		{"hours", `{{ humanizeDuration 3661 }}`, nil, "1h 1m 1s"},
		{"days", `{{ humanizeDuration 90061 }}`, nil, "1d 1h 1m 1s"},
		{"milliseconds", `{{ humanizeDuration 0.001 }}`, nil, "1ms"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := execTemplate(t, tt.template, tt.data)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToDate(t *testing.T) {
	t.Parallel()
	result := execTemplate(t, `{{ toDate "2006-01-02" "2024-06-15" | date "Jan 02, 2006" }}`, nil)
	assert.Equal(t, "Jun 15, 2024", result)
}

func TestMustToDate(t *testing.T) {
	t.Parallel()
	result := execTemplate(t, `{{ mustToDate "2006-01-02" "2024-12-25" | date "01/02/2006" }}`, nil)
	assert.Equal(t, "12/25/2024", result)
}

func TestMustToDate_Invalid(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New("test").Funcs(DefaultFuncs).Parse(`{{ mustToDate "2006-01-02" "not-a-date" }}`)
	require.NoError(t, err)
	err = tmpl.Execute(&writer{}, nil)
	assert.Error(t, err)
}

func TestSafeUrl(t *testing.T) {
	t.Parallel()
	result := execTemplate(t, `<a href="{{ safeUrl . }}">link</a>`, "https://example.com/path?q=1")
	assert.Equal(t, `<a href="https://example.com/path?q=1">link</a>`, result)
}

func TestUrlUnescape(t *testing.T) {
	t.Parallel()
	result := execTemplate(t, `{{ urlUnescape "hello%20world" }}`, nil)
	assert.Equal(t, "hello world", result)
}

func TestDict(t *testing.T) {
	t.Parallel()
	result := execTemplate(t, `{{ $d := dict "name" "alert1" "severity" "critical" }}{{ $d.name }}:{{ $d.severity }}`, nil)
	assert.Equal(t, "alert1:critical", result)
}

func TestDict_OddArgs(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New("test").Funcs(DefaultFuncs).Parse(`{{ dict "a" }}`)
	require.NoError(t, err)
	err = tmpl.Execute(&writer{}, nil)
	assert.Error(t, err)
}

func TestList(t *testing.T) {
	t.Parallel()
	result := execTemplate(t, `{{ $l := list "a" "b" "c" }}{{ index $l 1 }}`, nil)
	assert.Equal(t, "b", result)
}

func TestAppend(t *testing.T) {
	t.Parallel()
	result := execTemplate(t, `{{ $l := list "a" }}{{ $l2 := append $l "b" }}{{ len $l2 }}`, nil)
	assert.Equal(t, "2", result)
}

func TestPipeline_DateTzHumanize(t *testing.T) {
	t.Parallel()
	// Test a realistic pipeline: parse time → convert timezone → format
	result := execTemplate(t,
		`{{ toDate "2006-01-02T15:04" "2024-06-15T08:30" | tz "Asia/Shanghai" | date "2006-01-02 15:04 MST" }}`,
		nil,
	)
	assert.Equal(t, "2024-06-15 16:30 CST", result)
}
