package msgbus

import (
	"context"
)

func New(ctx context.Context, cfg *Configs) (MsgBus, error) {
	// base con Dsn we will use different msgbus,
	// use Nats.io by default
	return NewNats(ctx, cfg)
}

type MsgBus interface {
	Pub(ctx context.Context, event *Event) (*PubRes, error)
	Sub(ctx context.Context, sample *Event, queue string, fn func(ctx context.Context, event *Event) error) (func() error, error)

	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
}

type PubRes struct {
	Key string
}
