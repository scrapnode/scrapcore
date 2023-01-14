package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/scrapnode/scrapcore/msgbus/configs"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scrapcore/xmonitor"
	"go.uber.org/zap"
	"sync"
)

type Nats struct {
	Configs *configs.Configs
	Logger  *zap.SugaredLogger
	Monitor xmonitor.Monitor

	mu   sync.Mutex
	conn *nats.Conn
	jsc  nats.JetStreamContext
}

func New(ctx context.Context, cfg *configs.Configs) (*Nats, error) {
	logger := xlogger.FromContext(ctx).With("pkg", "scrapstream.msgbus.nats")
	monitor := xmonitor.FromContext(ctx)
	return &Nats{Configs: cfg, Logger: logger, Monitor: monitor}, nil
}
