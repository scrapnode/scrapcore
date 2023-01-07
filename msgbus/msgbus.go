package msgbus

import "context"

type SubscribeFn func(ctx context.Context, event *Event) error

type MsgBus interface {
	Pub(ctx context.Context, event *Event) (*PubRes, error)
	Sub(ctx context.Context, sample *Event, queue string, fn SubscribeFn) (func() error, error)

	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
}

type PubRes struct {
	Key string
}
