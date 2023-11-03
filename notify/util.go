package notify

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/template"
	"io"
)

// GetTemplateData creates the template data from the context and the alerts.
func GetTemplateData(ctx context.Context, tmpl *template.Template, alerts []*alert.Alert) *template.Data {
	recv, ok := ReceiverName(ctx)
	if !ok {
		log.Errorf("Missing receiver")
	}
	groupLabels, ok := GroupLabels(ctx)
	if !ok {
		log.Errorf("Missing group labels")
	}
	return tmpl.Data(recv, groupLabels, alerts...)
}

// TmplText is using monadic error handling in order to make string templating
// less verbose. Use with care as the final error checking is easily missed.
func TmplText(tmpl *template.Template, data *template.Data, err *error) func(string) string {
	return func(name string) (s string) {
		if *err != nil {
			return
		}
		s, *err = tmpl.ExecuteTextString(name, data)
		return s
	}
}

// TmplHTML is using monadic error handling in order to make string templating
// less verbose. Use with care as the final error checking is easily missed.
func TmplHTML(tmpl *template.Template, data *template.Data, err *error) func(string) string {
	return func(name string) (s string) {
		if *err != nil {
			return
		}
		s, *err = tmpl.ExecuteHTMLString(name, data)
		return s
	}
}

// Key is a string that can be hashed.
type Key string

// ExtractGroupKey gets the group key from the context.
func ExtractGroupKey(ctx context.Context) (Key, error) {
	key, ok := GroupKey(ctx)
	if !ok {
		return "", errors.New("group key missing")
	}
	return Key(key), nil
}

// Hash returns the sha256 for a group key as integrations may have
// maximum length requirements on deduplication keys.
func (k Key) Hash() string {
	h := sha256.New()
	// hash.Hash.Write never returns an error.
	//nolint: errcheck
	h.Write([]byte(string(k)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (k Key) String() string {
	return string(k)
}

// Retrier knows when to retry an HTTP request to a receiver. 2xx status codes
// are successful, anything else is a failure and only 5xx status codes should
// be retried.
type Retrier struct {
	// Function to return additional information in the error message.
	CustomDetailsFunc func(code int, body io.Reader) string
	// Additional HTTP status codes that should be retried.
	RetryCodes []int
}

// Check returns a boolean indicating whether the request should be retried
// and an optional error if the request has failed. If body is not nil, it will
// be included in the error message.
func (r *Retrier) Check(statusCode int, body io.Reader) (bool, error) {
	// 2xx responses are considered to be always successful.
	if statusCode/100 == 2 {
		return false, nil
	}

	// 5xx responses are considered to be always retried.
	retry := statusCode/100 == 5
	if !retry {
		for _, code := range r.RetryCodes {
			if code == statusCode {
				retry = true
				break
			}
		}
	}

	s := fmt.Sprintf("unexpected status code %v", statusCode)
	var details string
	if r.CustomDetailsFunc != nil {
		details = r.CustomDetailsFunc(statusCode, body)
	} else {
		details = readAll(body)
	}
	if details != "" {
		s = fmt.Sprintf("%s: %s", s, details)
	}
	return retry, errors.New(s)
}

func readAll(r io.Reader) string {
	if r == nil {
		return ""
	}
	bs, err := io.ReadAll(r)
	if err != nil {
		return ""
	}
	return string(bs)
}
