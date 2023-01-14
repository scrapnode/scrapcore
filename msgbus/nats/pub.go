package nats

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/msgbus/entity"
	"strings"
)

func (natsbus *Nats) Pub(ctx context.Context, event *entity.Event) (*msgbus.PubRes, error) {
	// @TODO: validator
	logger := natsbus.Logger.With("event_key", event.Key())
	event.SetMetadata(natsbus.Monitor.Propergator().Extract(ctx))

	msg, err := NewMsg(natsbus.Configs, event)
	if err != nil {
		logger.Errorw("could not construct Nats message from event", "error", err.Error())
		return nil, err
	}

	ack, err := natsbus.jsc.PublishMsg(msg)
	if err != nil {
		logger.Errorw("could not publish message to Nats", "error", err.Error())
		return nil, err
	}

	segments := []string{ack.Domain, ack.Stream, fmt.Sprint(ack.Sequence), event.Id}
	res := &msgbus.PubRes{Key: strings.Join(segments, "/")}

	logger.Debugw("published", "publish_key", res.Key)
	return res, nil
}
