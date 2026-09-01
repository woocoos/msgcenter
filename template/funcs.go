package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"time"

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
		return cases.Title(language.AmericanEnglish).String(text)
	},
	"trimSpace": strings.TrimSpace,
	"join": func(sep string, s []string) string {
		return strings.Join(s, sep)
	},
	"match": regexp.MatchString,
	"safeHtml": func(text string) tmplhtml.HTML {
		return tmplhtml.HTML(text)
	},
	"safeUrl": func(text string) tmplhtml.URL {
		return tmplhtml.URL(text)
	},
	"urlUnescape": url.QueryUnescape,
	"reReplaceAll": func(pattern, repl, text string) string {
		re := regexp.MustCompile(pattern)
		return re.ReplaceAllString(text, repl)
	},
	"stringSlice": func(s ...string) []string {
		return s
	},
	"markdown":       markdownEscapeString,
	"toJSON":         toJson,
	"now":            time.Now,
	"since":          time.Since,
	"date":           func(fmt string, t time.Time) string { return t.Format(fmt) },
	"tz":             func(name string, t time.Time) (time.Time, error) { loc, err := time.LoadLocation(name); if err != nil { return time.Time{}, err }; return t.In(loc), nil },
	"toDate":         func(layout, s string) time.Time { t, _ := time.ParseInLocation(layout, s, time.UTC); return t },
	"mustToDate":     func(layout, s string) (time.Time, error) { return time.ParseInLocation(layout, s, time.UTC) },
	"humanizeDuration": humanizeDuration,
	"list": func(args ...any) ([]any, error) {
		if args == nil {
			return []any{}, nil
		}
		return args, nil
	},
	"append": func(slice []any, args ...any) []any {
		return append(slice, args...)
	},
	"dict": func(values ...any) (map[string]any, error) {
		if len(values)%2 != 0 {
			return nil, fmt.Errorf("dict requires an even number of arguments")
		}
		res := make(map[string]any, len(values)/2)
		for i := 0; i < len(values); i += 2 {
			key, ok := values[i].(string)
			if !ok {
				return nil, fmt.Errorf("dict keys must be strings")
			}
			res[key] = values[i+1]
		}
		return res, nil
	},
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

func convertToFloat(i any) (float64, error) {
	switch v := i.(type) {
	case float64:
		return v, nil
	case string:
		return strconv.ParseFloat(v, 64)
	case int:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case time.Duration:
		return v.Seconds(), nil
	default:
		return 0, fmt.Errorf("can't convert %T to float", v)
	}
}

func humanizeDuration(i any) (string, error) {
	v, err := convertToFloat(i)
	if err != nil {
		return "", err
	}
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return fmt.Sprintf("%.4g", v), nil
	}
	if v == 0 {
		return "0s", nil
	}
	if math.Abs(v) >= 1 {
		sign := ""
		if v < 0 {
			sign = "-"
			v = -v
		}
		dur := int64(v)
		seconds := dur % 60
		minutes := (dur / 60) % 60
		hours := (dur / 3600) % 24
		days := dur / 86400
		switch {
		case days != 0:
			return fmt.Sprintf("%s%dd %dh %dm %ds", sign, days, hours, minutes, seconds), nil
		case hours != 0:
			return fmt.Sprintf("%s%dh %dm %ds", sign, hours, minutes, seconds), nil
		case minutes != 0:
			return fmt.Sprintf("%s%dm %ds", sign, minutes, seconds), nil
		default:
			return fmt.Sprintf("%s%.4gs", sign, v), nil
		}
	}
	prefix := ""
	for _, p := range []string{"m", "u", "n", "p", "f", "a", "z", "y"} {
		if math.Abs(v) >= 1 {
			break
		}
		prefix = p
		v *= 1000
	}
	return fmt.Sprintf("%.4g%ss", v, prefix), nil
}
