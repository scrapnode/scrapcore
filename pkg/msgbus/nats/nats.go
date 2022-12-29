package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/scrapnode/scrapcore/pkg/msgbus/configs"
	"go.uber.org/zap"
	"sync"
)

type Nats struct {
	Configs *configs.Configs
	Logger  *zap.SugaredLogger

	mu   sync.Mutex
	conn *nats.Conn
	jsc  nats.JetStreamContext
}
