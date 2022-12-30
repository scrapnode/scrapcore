package sql

import (
	"context"
	"github.com/scrapnode/scrapcore/database/configs"
	"github.com/scrapnode/scrapcore/xlogger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

type SQL struct {
	Configs *configs.Configs
	Logger  *zap.SugaredLogger
	Conn    *gorm.DB

	mu sync.Mutex
}

func New(ctx context.Context, cfg *configs.Configs) (*SQL, error) {
	logger := xlogger.FromContext(ctx).With("pkg", "database.sql")
	return &SQL{Configs: cfg, Logger: logger}, nil
}
