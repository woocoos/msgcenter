package template

import (
	"bytes"
	"encoding/json"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	tmplhtml "html/template"
	"regexp"
	"strings"
)

var (
	isMarkdownSpecial [128]bool
)

var DefaultFuncs = map[string]any{
	"toUpper": strings.ToUpper,
	"toLower": strings.ToLower,
	"title": func(text string) string {
		// Casers should not be shared between goroutines, instead
		// create a new caser each time this function is called.
		return cases.Title(language.AmericanEnglish).String(text)
	},
	"trimSpace": strings.TrimSpace,
	// join is equal to strings.Join but inverts the argument order
	// for easier pipelining in templates.
	"join": func(sep string, s []string) string {
		return strings.Join(s, sep)
	},
	"match": regexp.MatchString,
	"safeHtml": func(text string) tmplhtml.HTML {
		return tmplhtml.HTML(text)
	},
	"reReplaceAll": func(pattern, repl, text string) string {
		re := regexp.MustCompile(pattern)
		return re.ReplaceAllString(text, repl)
	},
	"stringSlice": func(s ...string) []string {
		return s
	},
	"markdown": markdownEscapeString,
	"toJSON":   toJson,
}

func init() {
	for _, c := range "_*`" {
		isMarkdownSpecial[c] = true
	}
}

func markdownEscapeString(s string) string {
	b := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(b)

	for _, c := range s {
		if c < 128 && isMarkdownSpecial[c] {
			buf.WriteByte('\\')
		}
		buf.WriteRune(c)
	}
	return buf.String()
}

func toJson(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
