package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/samber/lo"
	"github.com/scrapnode/scrapcore/msgbus/configs"
	"github.com/scrapnode/scrapcore/msgbus/entity"
	"strconv"
	"strings"
)

func NewMsg(cfg *configs.Configs, event *entity.Event) (*nats.Msg, error) {
	msg := nats.NewMsg(NewSubject(cfg, event))
	msg.Data = event.Data

	msg.Header.Set("Nats-Msg-Id", event.Id)
	// with metadata
	msg.Header.Set(entity.METAKEY_WORKSPACE, event.Workspace)
	msg.Header.Set(entity.METAKEY_APP, event.App)
	msg.Header.Set(entity.METAKEY_TYPE, event.Type)
	msg.Header.Set(entity.METAKEY_ID, event.Id)
	msg.Header.Set(entity.METAKEY_TIMESTAMPS, fmt.Sprint(event.Timestamps))

	if event.Metadata != nil {
		for key, value := range event.Metadata {
			if lo.Contains(entity.METAKEY_RESERVE, key) {
				continue
			}
			if strings.HasPrefix(key, "Nats") {
				continue
			}

			msg.Header.Set(key, value)
		}
	}

	return msg, nil
}

func NewEvent(msg *nats.Msg) (*entity.Event, error) {
	event := &entity.Event{
		Workspace: msg.Header.Get(entity.METAKEY_WORKSPACE),
		App:       msg.Header.Get(entity.METAKEY_APP),
		Type:      msg.Header.Get(entity.METAKEY_TYPE),
		Id:        msg.Header.Get(entity.METAKEY_ID),
		Data:      msg.Data,
		Metadata:  map[string]string{},
	}

	ts, err := strconv.ParseInt(msg.Header.Get(entity.METAKEY_TIMESTAMPS), 10, 64)
	if err != nil {
		return nil, err
	}
	event.Timestamps = ts

	for key, value := range msg.Header {
		if lo.Contains(entity.METAKEY_RESERVE, key) {
			continue
		}
		if strings.HasPrefix(key, "Nats") {
			continue
		}

		event.Metadata[key] = value[0]
	}

	return event, nil
}
