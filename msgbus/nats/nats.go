package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/msgbus/configs"
	"github.com/scrapnode/scrapcore/xlogger"
	"sync"
)

type Nats struct {
	Configs *configs.Configs
	Logger  *xlogger.Logger

	mu   sync.Mutex
	conn *nats.Conn
	jsc  nats.JetStreamContext
}

func New(ctx context.Context, cfg *configs.Configs) (msgbus.MsgBus, error) {
	logger := xlogger.FromContext(ctx).With("pkg", "scrapstream.msgbus.nats")
	return &Nats{Configs: cfg, Logger: logger}, nil
}
