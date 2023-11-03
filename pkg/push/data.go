package push

import (
	"encoding/json"
	"fmt"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/pkg/label"
)

// Data push protocol
type Data struct {
	Topic    string   `json:"topic"`
	Audience Audience `json:"audience"`
	Message  Message  `json:"message"`
}

// Message main body
type Message struct {
	Title   string             `json:"title"`
	Content string             `json:"content"`
	Format  msgtemplate.Format `json:"format"`
	Extras  label.LabelSet     `json:"extras"`
}

// Audience is the targets that can receive the message.
type Audience struct {
	// AppCode app to receive
	AppCode string `json:"appCode"`
	// UserIDs users id list
	UserIDs []int `json:"userIds"`
	// DeviceIDs device list
	DeviceIDs []string `json:"deviceIDs"`
}

func (d *Data) Validate() error {
	if d.Topic == "" {
		return fmt.Errorf("topic missing")
	}

	return nil
}

func Marshal(data any) ([]byte, error) {
	return json.Marshal(data)
}

func Unmarshal(v []byte) (*Data, error) {
	var d Data
	err := json.Unmarshal(v, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
