package template

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	tmplhtml "html/template"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	tmpltext "text/template"
	"text/template/parse"
)

var (
	//go:embed *
	templateDir embed.FS
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
}

// Template bundles a text and a html template instance.
type Template struct {
	text *tmpltext.Template
	html *tmplhtml.Template

	ExternalURL *url.URL
}

// Option is generic modifier of the text and html templates used by a Template.
type Option func(text *tmpltext.Template, html *tmplhtml.Template)

// New returns a new Template with the DefaultFuncs added. The DefaultFuncs
// have precedence over any added custom functions. Options allow customization
// of the text and html templates in given order.
func New(options ...Option) (*Template, error) {
	t := &Template{
		text: tmpltext.New("").Option("missingkey=zero"),
		html: tmplhtml.New("").Option("missingkey=zero"),
	}

	for _, o := range options {
		o(t.text, t.html)
	}
	t.text.Funcs(DefaultFuncs)
	t.html.Funcs(DefaultFuncs)
	MustParse(t.ParseFS(templateDir, "*.tmpl"))
	return t, nil
}

// Parse parses text as a template body for t.
func (t *Template) Parse(text string) (*Template, error) {
	if _, err := t.text.Parse(text); err != nil {
		return nil, err
	}
	if _, err := t.html.Parse(text); err != nil {
		return nil, err
	}
	return t, nil
}

// ParseFiles parses a list of files as templates and associate them with t.
// Each file can be a standalone template.
func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
	if _, err := t.text.ParseFiles(filenames...); err != nil {
		return nil, err
	}
	if _, err := t.html.ParseFiles(filenames...); err != nil {
		return nil, err
	}
	return t, nil
}

// ParseGlob parses the files that match the given pattern as templates and
// associate them with t.
func (t *Template) ParseGlob(pattern string) (*Template, error) {
	if _, err := t.text.ParseGlob(pattern); err != nil {
		return nil, err
	}
	if _, err := t.html.ParseGlob(pattern); err != nil {
		return nil, err
	}
	return t, nil
}

// ParseDir walks on the given dir path and parses the given matches with aren't Go files.
func (t *Template) ParseDir(path string) (*Template, error) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walk path %s: %w", path, err)
		}
		if info.IsDir() || strings.HasSuffix(path, ".go") {
			return nil
		}
		_, err = t.ParseFiles(path)
		return err
	})
	return t, err
}

// ParseFS is like ParseFiles or ParseGlob but reads from the file system fsys
// instead of the host operating system's file system.
func (t *Template) ParseFS(fsys fs.FS, patterns ...string) (*Template, error) {
	if _, err := t.text.ParseFS(fsys, patterns...); err != nil {
		return nil, err
	}
	if _, err := t.html.ParseFS(fsys, patterns...); err != nil {
		return nil, err
	}
	return t, nil
}

// AddParseTree adds the given parse tree to the template.
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error) {
	if _, err := t.text.AddParseTree(name, tree); err != nil {
		return nil, err
	}
	if _, err := t.html.AddParseTree(name, tree); err != nil {
		return nil, err
	}
	return t, nil
}

// MustParse is a helper that wraps a call to a function returning (*Template, error)
// and panics if the error is non-nil.
func MustParse(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}

// ExecuteTextString needs a meaningful doc comment (TODO(fabxc)).
func (t *Template) ExecuteTextString(text string, data interface{}) (string, error) {
	if text == "" {
		return "", nil
	}
	tmpl, err := t.text.Clone()
	if err != nil {
		return "", err
	}
	tmpl, err = tmpl.New("").Option("missingkey=zero").Parse(text)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	return buf.String(), err
}

// ExecuteHTMLString needs a meaningful doc comment (TODO(fabxc)).
func (t *Template) ExecuteHTMLString(html string, data interface{}) (string, error) {
	if html == "" {
		return "", nil
	}
	tmpl, err := t.html.Clone()
	if err != nil {
		return "", err
	}
	tmpl, err = tmpl.New("").Option("missingkey=zero").Parse(html)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	return buf.String(), err
}

// Data assembles data for template expansion.
func (t *Template) Data(recv string, groupLabels label.LabelSet, alerts ...*alert.Alert) *Data {
	data := &Data{
		Receiver:          regexp.QuoteMeta(recv),
		Status:            string(alert.Alerts(alerts).Status()),
		Alerts:            make([]alert.Alert, 0, len(alerts)),
		GroupLabels:       KV{},
		CommonLabels:      KV{},
		CommonAnnotations: KV{},
	}
	if t.ExternalURL != nil {
		data.ExternalURL = t.ExternalURL.String()
	}

	for _, a := range alerts {
		data.Alerts = append(data.Alerts, *a.Clone())
	}

	for k, v := range groupLabels {
		data.GroupLabels[string(k)] = string(v)
	}

	if len(alerts) >= 1 {
		var (
			commonLabels      = alerts[0].Labels.Clone()
			commonAnnotations = alerts[0].Annotations.Clone()
		)
		for _, a := range alerts[1:] {
			if len(commonLabels) == 0 && len(commonAnnotations) == 0 {
				break
			}
			for ln, lv := range commonLabels {
				if a.Labels[ln] != lv {
					delete(commonLabels, ln)
				}
			}
			for an, av := range commonAnnotations {
				if a.Annotations[an] != av {
					delete(commonAnnotations, an)
				}
			}
		}
		for k, v := range commonLabels {
			data.CommonLabels[string(k)] = string(v)
		}
		for k, v := range commonAnnotations {
			data.CommonAnnotations[string(k)] = string(v)
		}
	}

	return data
}
