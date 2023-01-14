package msgbus

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scrapcore/xmonitor"
	"go.uber.org/zap"
	"sync"
)

func NewNats(ctx context.Context, cfg *Configs) (*Nats, error) {
	logger := xlogger.FromContext(ctx).With("pkg", "scrapstream.msgbus.nats")
	monitor := xmonitor.FromContext(ctx)
	return &Nats{cfg: cfg, logger: logger, monitor: monitor}, nil
}

type Nats struct {
	cfg     *Configs
	logger  *zap.SugaredLogger
	monitor xmonitor.Monitor

	mu   sync.Mutex
	conn *nats.Conn
	jsc  nats.JetStreamContext
}
