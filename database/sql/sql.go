package sql

import (
	"context"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/xlogger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

type SQL struct {
	Configs *database.Configs
	Logger  *zap.SugaredLogger
	Conn    *gorm.DB

	mu sync.Mutex
}

func New(ctx context.Context, cfg *database.Configs) (*SQL, error) {
	logger := xlogger.FromContext(ctx).With("pkg", "database.sql")
	return &SQL{Configs: cfg, Logger: logger}, nil
}
