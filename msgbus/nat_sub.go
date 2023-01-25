package msgbus

import (
	"context"
	"github.com/nats-io/nats.go"
	"time"
)

func (natsbus *Nats) Sub(ctx context.Context, sample *Event, queue string, fn SubscribeFn) (func() error, error) {
	subject := NatsSubject(natsbus.cfg, sample)
	opts := []nats.SubOpt{
		nats.DeliverNew(),
		nats.AckExplicit(),
		nats.MaxDeliver(natsbus.cfg.MaxRetry + 1),
		nats.BackOff(NewBackoff(natsbus.cfg.MaxRetry)),
	}

	sub, err := natsbus.jsc.QueueSubscribe(subject, queue, natsbus.UseSub(fn), opts...)
	// by default the consumer that is created by QueueSubscribe will be there forever (set durable to TRUE)
	if err != nil {
		return func() error { return nil }, err
	}

	natsbus.logger.Debugw("subscribed", "subject", subject, "queue", queue)
	return func() error { return sub.Drain() }, err
}

func (natsbus *Nats) UseSub(fn func(ctx context.Context, event *Event) error) nats.MsgHandler {
	delay := 5 * time.Second
	backoff := NewBackoff(natsbus.cfg.MaxRetry)
	if len(backoff) > 0 {
		delay = backoff[0]
	}

	return func(msg *nats.Msg) {
		event, err := NatsMsg2Event(msg)
		if err != nil {
			natsbus.logger.Errorw("could not parse event from message", "error", err.Error())
			if err := msg.Ack(); err != nil {
				natsbus.logger.Errorw("ack was failed", "error", err.Error())
			}
			return
		}

		logger := natsbus.logger.With("event_key", event.Key())
		logger.Debug("got event")

		ctx := natsbus.monitor.Propergator().Inject(context.Background(), event.Metadata)
		// handler of subscription must handle all the error, if it returns any error, we will trigger retry
		if err := fn(ctx, event); err != nil {
			logger.Errorw("could not handle event", "error", err.Error())
			// nats.BackOff does not work with QueueSubscribe, so we will fall back to first value of nats.BackOff
			// we cannot retry by ourselves with some hack of set headers and Nak it
			if err := msg.NakWithDelay(delay); err != nil {
				logger.Errorw("nak was failed", "error", err.Error())
				return
			}

			logger.Infow("nak was succesful", "error", err.Error())
			return
		}

		if err := msg.Ack(); err != nil {
			logger.Errorw("ack was failed", "error", err.Error())
			return
		}

		logger.Debug("processed event")
	}
}
