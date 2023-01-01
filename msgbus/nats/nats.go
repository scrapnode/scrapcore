package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/xlogger"
	"go.uber.org/zap"
	"sync"
)

type Nats struct {
	Configs *msgbus.Configs
	Logger  *zap.SugaredLogger

	mu   sync.Mutex
	conn *nats.Conn
	jsc  nats.JetStreamContext
}

func New(ctx context.Context, cfg *msgbus.Configs) (msgbus.MsgBus, error) {
	logger := xlogger.FromContext(ctx).With("pkg", "scrapstream.msgbus.nats")
	return &Nats{Configs: cfg, Logger: logger}, nil
}
