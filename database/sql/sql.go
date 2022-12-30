package sql

import (
	"context"
	"github.com/scrapnode/scrapcore/database/configs"
	"github.com/scrapnode/scrapcore/xlogger"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/url"
)

func New(ctx context.Context, cfg *configs.Configs) (*gorm.DB, error) {
	dialector, err := Dial(cfg.Dsn)
	if err != nil {
		return nil, err
	}

	logger := &Logger{zap: xlogger.FromContext(ctx).With("package", "database.sql")}
	return gorm.Open(dialector, &gorm.Config{
		Logger: logger,
	})
}

// Dial support both SQLite & PostgreSQL
func Dial(dsn string) (gorm.Dialector, error) {
	uri, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	if uri.Scheme == "sqlite" {
		return sqlite.Open(uri.Host + uri.Path + uri.RawQuery), nil
	}

	return postgres.Open(dsn), nil
}
