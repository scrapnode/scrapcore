package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/samber/lo"
	"github.com/scrapnode/scrapcore/msgbus"
	"strconv"
	"strings"
)

func NewMsg(cfg *msgbus.Configs, event *msgbus.Event) (*nats.Msg, error) {
	msg := nats.NewMsg(NewSubject(cfg, event))
	msg.Data = []byte(event.Data)

	msg.Header.Set("Nats-Msg-Id", event.Id)
	// with metadata
	msg.Header.Set(msgbus.METAKEY_WORKSPACE, event.Workspace)
	msg.Header.Set(msgbus.METAKEY_APP, event.App)
	msg.Header.Set(msgbus.METAKEY_TYPE, event.Type)
	msg.Header.Set(msgbus.METAKEY_ID, event.Id)
	msg.Header.Set(msgbus.METAKEY_TIMESTAMPS, fmt.Sprint(event.Timestamps))

	if event.Metadata != nil {
		for key, value := range event.Metadata {
			if strings.HasPrefix(key, msgbus.METAKEY_PREFIX) {
				msg.Header.Set(key, value)
			}
		}
	}

	return msg, nil
}

func NewEvent(msg *nats.Msg) (*msgbus.Event, error) {
	event := &msgbus.Event{
		Workspace: msg.Header.Get(msgbus.METAKEY_WORKSPACE),
		App:       msg.Header.Get(msgbus.METAKEY_APP),
		Type:      msg.Header.Get(msgbus.METAKEY_TYPE),
		Id:        msg.Header.Get(msgbus.METAKEY_ID),
		Data:      msg.Data,
		Metadata:  map[string]string{},
	}

	ts, err := strconv.ParseInt(msg.Header.Get(msgbus.METAKEY_TIMESTAMPS), 10, 64)
	if err != nil {
		return nil, err
	}
	event.Timestamps = ts

	for key, value := range msg.Header {
		if !strings.HasPrefix(key, msgbus.METAKEY_PREFIX) {
			continue
		}
		if lo.Contains(msgbus.METAKEY_RESERVE, key) {
			continue
		}

		event.Metadata[key] = value[0]
	}

	return event, nil
}
