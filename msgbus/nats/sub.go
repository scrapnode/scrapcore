package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	msgbus2 "github.com/scrapnode/scrapcore/msgbus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"time"
)

func (natsbus *Nats) Sub(ctx context.Context, sample *msgbus2.Event, queue string, fn msgbus2.SubscribeFn) (func() error, error) {
	subject := NewSubject(natsbus.Configs, sample)
	opts := []nats.SubOpt{
		nats.DeliverNew(),
		nats.AckExplicit(),
		nats.MaxDeliver(natsbus.Configs.MaxRetry + 1),
		nats.BackOff(NewBackoff(natsbus.Configs.MaxRetry)),
	}

	sub, err := natsbus.jsc.QueueSubscribe(subject, queue, natsbus.UseSub(fn), opts...)
	// by default the consumer that is created by QueueSubscribe will be there forever (set durable to TRUE)
	if err != nil {
		return func() error { return nil }, err
	}

	natsbus.Logger.Debugw("subscribed", "subject", subject, "queue", queue)
	return func() error { return sub.Drain() }, err
}

func (natsbus *Nats) UseSub(fn msgbus2.SubscribeFn) nats.MsgHandler {
	delay := 5 * time.Second
	backoff := NewBackoff(natsbus.Configs.MaxRetry)
	if len(backoff) > 0 {
		delay = backoff[0]
	}

	return func(msg *nats.Msg) {
		event, err := NewEvent(msg)
		if err != nil {
			natsbus.Logger.Errorw("could not parse event from message", "error", err.Error())
			if err := msg.Ack(); err != nil {
				natsbus.Logger.Errorw("ack was failed", "error", err.Error())
			}
			return
		}
		ctx := otel.GetTextMapPropagator().Extract(context.Background(), propagation.MapCarrier(event.Metadata))

		logger := natsbus.Logger.With("event_key", event.Key())
		logger.Debug("got event")

		// handler of subscription must handle all the error, if it returns any error, we will trigger retry
		if err := fn(ctx, event); err != nil {
			logger.Errorw("could not handle event", "error", err.Error())
			// nats.BackOff does not work with QueueSubscribe, so we will fall back to first value of nats.BackOff
			// we cannot retry by ourselves with some hack of set headers and Nak it
			if err := msg.NakWithDelay(delay); err != nil {
				logger.Errorw("nak was failed", "error", err.Error())
				return
			}

			logger.Errorw("nak was succesful", "error", err.Error())
			return
		}

		if err := msg.Ack(); err != nil {
			logger.Errorw("ack was failed", "error", err.Error())
			return
		}

		logger.Debug("processed event")
	}
}
