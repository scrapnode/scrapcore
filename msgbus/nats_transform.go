package msgbus

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func Event2NatsMsg(cfg *Configs, event *Event) (*nats.Msg, error) {
	msg := nats.NewMsg(NatsSubject(cfg, event))
	msg.Data = event.Data

	msg.Header.Set("Nats-Msg-Id", event.Id)
	// with metadata
	msg.Header.Set(METAKEY_WORKSPACE, event.Workspace)
	msg.Header.Set(METAKEY_APP, event.App)
	msg.Header.Set(METAKEY_TYPE, event.Type)
	msg.Header.Set(METAKEY_ID, event.Id)
	msg.Header.Set(METAKEY_TIMESTAMPS, fmt.Sprint(event.Timestamps))

	if event.Metadata != nil {
		for key, value := range event.Metadata {
			if lo.Contains(METAKEY_RESERVE, key) {
				continue
			}
			if strings.HasPrefix(strings.ToLower(key), "nats") {
				continue
			}

			msg.Header.Set(key, value)
		}
	}

	return msg, nil
}

func NatsMsg2Event(msg *nats.Msg) (*Event, error) {
	event := &Event{
		Workspace: msg.Header.Get(METAKEY_WORKSPACE),
		App:       msg.Header.Get(METAKEY_APP),
		Type:      msg.Header.Get(METAKEY_TYPE),
		Id:        msg.Header.Get(METAKEY_ID),
		Data:      msg.Data,
		Metadata:  map[string]string{},
	}

	ts, err := strconv.ParseInt(msg.Header.Get(METAKEY_TIMESTAMPS), 10, 64)
	if err != nil {
		return nil, err
	}
	event.Timestamps = ts

	for key, value := range msg.Header {
		if lo.Contains(METAKEY_RESERVE, key) {
			continue
		}
		if strings.HasPrefix(strings.ToLower(key), "nats") {
			continue
		}

		event.Metadata[key] = value[0]
	}

	return event, nil
}
