package msgbus

import "context"

type SubscribeFn func(event *Event) error

type MsgBus interface {
	Pub(ctx context.Context, event *Event) (*PubRes, error)
	Sub(ctx context.Context, sample *Event, queue string, fn SubscribeFn) error

	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
}

type PubRes struct {
	Key string
}
