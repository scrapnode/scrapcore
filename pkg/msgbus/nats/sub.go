package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	msgbus2 "github.com/scrapnode/scrapcore/pkg/msgbus"
	"time"
)

func (natsbus *Nats) Sub(sample *msgbus2.Event, queue string, fn msgbus2.SubscribeFn) error {
	subject := NewSubject(natsbus.Configs, sample)
	opts := []nats.SubOpt{
		nats.DeliverNew(),
		nats.AckExplicit(),
		nats.MaxDeliver(natsbus.Configs.MaxRetry + 1),
		nats.BackOff(NewBackoff(natsbus.Configs.MaxRetry)),
	}

	// by default the consumer that is created by QueueSubscribe will be there forever (set durable to TRUE)
	if _, err := natsbus.jsc.QueueSubscribe(subject, queue, natsbus.UseSub(fn), opts...); err != nil {
		return err
	}

	natsbus.Logger.Debugw("subscribed", "subject", subject, "queue", queue)
	return nil
}

func (natsbus *Nats) UseSub(fn msgbus2.SubscribeFn) nats.MsgHandler {
	delay := 5 * time.Second
	backoff := NewBackoff(natsbus.Configs.MaxRetry)
	if len(backoff) > 0 {
		delay = backoff[0]
	}

	return func(msg *nats.Msg) {
		ctx := context.Background()

		event, err := NewEvent(msg)
		if err != nil {
			natsbus.Logger.Errorw("could not parse event from message", "error", err.Error())
			if err := msg.Ack(nats.Context(ctx)); err != nil {
				natsbus.Logger.Errorw("ack was failed", "error", err.Error())
			}
			return
		}
		logger := natsbus.Logger.With("event_key", event.Key())
		logger.Debug("got event")

		// handler of subcription must handle all of the error, if it returns any error, we will trigger retry
		if err := fn(event); err != nil {
			logger.Errorw("could not handle event", "error", err.Error())
			// nats.BackOff does not work with QueueSubscribe so we will fallback to first value of nats.BackOff
			// we cannot retry by ourself with some hack of set headers and Nak it
			if err := msg.NakWithDelay(delay); err != nil {
				logger.Errorw("nak was failed", "error", err.Error())
				return
			}

			logger.Errorw("nak was succesful", "error", err.Error())
			return
		}

		if err := msg.Ack(nats.Context(ctx)); err != nil {
			logger.Errorw("ack was failed", "error", err.Error())
			return
		}

		logger.Debug("processed event")
	}
}
