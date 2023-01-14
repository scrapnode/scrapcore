package msgbus

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus/configs"
	"github.com/scrapnode/scrapcore/msgbus/entity"
	"github.com/scrapnode/scrapcore/msgbus/nats"
)

func New(ctx context.Context, cfg *configs.Configs) (MsgBus, error) {
	// base con Dsn we will use different msgbus,
	// use Nats.io by default
	return nats.New(ctx, cfg)
}

type MsgBus interface {
	Pub(ctx context.Context, event *entity.Event) (*PubRes, error)
	Sub(ctx context.Context, sample *entity.Event, queue string, fn func(ctx context.Context, event *entity.Event) error) (func() error, error)

	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
}

type PubRes struct {
	Key string
}
