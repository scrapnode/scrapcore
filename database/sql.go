package database

import (
	"context"
	"github.com/scrapnode/scrapcore/xlogger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

func NewSQL(ctx context.Context, cfg *Configs) (*SQL, error) {
	logger := xlogger.FromContext(ctx).With("pkg", "database.sql")
	return &SQL{cfg: cfg, logger: logger}, nil
}

type SQL struct {
	cfg    *Configs
	logger *zap.SugaredLogger

	conn *gorm.DB
	mu   sync.Mutex
}
