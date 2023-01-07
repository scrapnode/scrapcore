package nats

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/msgbus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"strings"
)

func (natsbus *Nats) Pub(ctx context.Context, event *msgbus.Event) (*msgbus.PubRes, error) {
	// @TODO: validator
	logger := natsbus.Logger.With("event_key", event.Key())

	if event.Metadata != nil {
		otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(event.Metadata))
	}

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
