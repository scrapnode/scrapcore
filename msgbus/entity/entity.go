package entity

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/scrapnode/scrapcore/utils"
	"strings"
)

var METAKEY_PREFIX = "x-scrapevent"
var (
	METAKEY_ID         = fmt.Sprintf("%s-id", METAKEY_PREFIX)
	METAKEY_WORKSPACE  = fmt.Sprintf("%s-workspace", METAKEY_PREFIX)
	METAKEY_APP        = fmt.Sprintf("%s-app", METAKEY_PREFIX)
	METAKEY_TYPE       = fmt.Sprintf("%s-type", METAKEY_PREFIX)
	METAKEY_TIMESTAMPS = fmt.Sprintf("%s-timestamps", METAKEY_PREFIX)
)
var METAKEY_RESERVE = []string{
	METAKEY_ID,
	METAKEY_WORKSPACE,
	METAKEY_APP,
	METAKEY_TYPE,
	METAKEY_TIMESTAMPS,
}

type Event struct {
	Workspace string `json:"workspace"`
	App       string `json:"app"`
	Type      string `json:"type"`

	Id         string            `json:"id"`
	Timestamps int64             `json:"timestamps"`
	Data       []byte            `json:"data"`
	Metadata   map[string]string `json:"metadata"`
}

func (event *Event) SetId() error {
	// only set data if it wasn't set yet
	if event.Id != "" {
		return errors.New("msgbus.entity: id has set already")
	}

	event.Id = utils.NewId("event")
	return nil
}

func (event *Event) SetData(data interface{}) error {
	if event.Data != nil {
		return errors.New("msgbus.entity: data has set already")
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	event.Data = bytes
	return nil
}

func (event *Event) ScanData(dest interface{}) error {
	if event.Data == nil {
		return errors.New("msgbus.entity: data is empty")
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(event.Data, dest)
}

func (event *Event) SetMetadata(meta map[string]string) {
	if event.Metadata == nil {
		event.Metadata = map[string]string{}
	}

	if meta == nil {
		return
	}

	for key, value := range meta {
		event.Metadata[key] = value
	}
}

func (event *Event) GetMetadata() map[string]string {
	if event.Metadata == nil {
		event.Metadata = map[string]string{}
	}

	return event.Metadata
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
