package notify

import (
	"context"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/template"
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
