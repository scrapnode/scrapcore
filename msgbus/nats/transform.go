package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/msgbus/configs"
	"strconv"
)

func NewMsg(cfg *configs.Configs, event *msgbus.Event) (*nats.Msg, error) {
	msg := nats.NewMsg(NewSubject(cfg, event))
	msg.Data = []byte(event.Data)

	msg.Header.Set("Nats-Msg-Id", event.Id)
	// with metadata
	msg.Header.Set(msgbus.METAKEY_EVENT_WORKSPACE, event.Workspace)
	msg.Header.Set(msgbus.METAKEY_EVENT_APP, event.App)
	msg.Header.Set(msgbus.METAKEY_EVENT_TYPE, event.Type)
	msg.Header.Set(msgbus.METAKEY_EVENT_ID, event.Id)
	msg.Header.Set(msgbus.METAKEY_EVENT_TIMESTAMPS, fmt.Sprint(event.Timestamps))

	return msg, nil
}

func NewEvent(msg *nats.Msg) (*msgbus.Event, error) {
	event := &msgbus.Event{
		Workspace: msg.Header.Get(msgbus.METAKEY_EVENT_WORKSPACE),
		App:       msg.Header.Get(msgbus.METAKEY_EVENT_APP),
		Type:      msg.Header.Get(msgbus.METAKEY_EVENT_TYPE),
		Id:        msg.Header.Get(msgbus.METAKEY_EVENT_ID),
		Data:      msg.Data,
	}

	ts, err := strconv.ParseInt(msg.Header.Get(msgbus.METAKEY_EVENT_TIMESTAMPS), 10, 64)
	if err != nil {
		return nil, err
	}
	event.Timestamps = ts

	return event, nil
}
