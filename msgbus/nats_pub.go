package msgbus

import (
	"context"
	"fmt"
	"strings"
)

func (natsbus *Nats) Pub(ctx context.Context, event *Event) (*PubRes, error) {
	// @TODO: validator
	logger := natsbus.logger.With("event_key", event.Key())
	event.SetMetadata(natsbus.monitor.Propergator().Extract(ctx))

	msg, err := Event2NatsMsg(natsbus.cfg, event)
	if err != nil {
		logger.Errorw("msgbus.nats: could not construct message from event", "error", err.Error())
		return nil, err
	}

	ack, err := natsbus.jsc.PublishMsg(msg)
	if err != nil {
		logger.Errorw("msgbus.nats: could not publish message to Nats", "error", err.Error())
		return nil, err
	}

	segments := []string{ack.Domain, ack.Stream, fmt.Sprint(ack.Sequence), event.Id}
	res := &PubRes{Key: strings.Join(segments, "/")}

	logger.Debugw("msgbus.nats: published", "publish_key", res.Key)
	return res, nil
}
