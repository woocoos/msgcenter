package webhook

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/woocoos/msgcenter/template"
	"testing"
)

func TestMessageJson(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		msg := Message{
			Version: "4",
			Data: &template.Data{
				Receiver: "test",
				Status:   "firing",
				Alerts: template.Alerts{
					{
						Status: "firing",
					},
				},
			},
			GroupKey: "test",
		}
		bs, err := json.Marshal(msg)
		assert.NoError(t, err)
		t.Log(string(bs))
	})
}
