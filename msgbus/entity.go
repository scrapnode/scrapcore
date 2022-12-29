package msgbus

import (
	"encoding/json"
	"github.com/scrapnode/scrapcore/utils"
	"strings"
)

var (
	METAKEY_EVENT_ID         = "X-ScrapCore-Event-Id"
	METAKEY_EVENT_WORKSPACE  = "X-ScrapCore-Event-Workspace"
	METAKEY_EVENT_APP        = "X-ScrapCore-Event-App"
	METAKEY_EVENT_TYPE       = "X-ScrapCore-Event-Type"
	METAKEY_EVENT_TIMESTAMPS = "X-ScrapCore-Event-Timestamps"
)

type Event struct {
	Workspace string `json:"workspace"`
	App       string `json:"app"`
	Type      string `json:"type"`

	Id         string `json:"id"`
	Timestamps int64  `json:"timestamps"`
	Data       []byte `json:"data"`
}

func (event *Event) SetId() error {
	// only set data if it wasn't set yet
	if event.Id != "" {
		return ErrEventIdWasSet
	}

	event.Id = utils.NewId("event")
	return nil
}

func (event *Event) SetData(data interface{}) error {
	if event.Data != nil {
		return ErrEventDataWasSet
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	event.Data = bytes
	return nil
}

func (event *Event) Key() string {
	keys := []string{
		event.Workspace,
		event.App,
		event.Type,
		event.Id,
	}
	return strings.Join(keys, "/")
}
