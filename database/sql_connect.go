package database

import (
	"context"
	"github.com/scrapnode/scrapcore/xlogger"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/url"
	"strings"
)

func (db *SQL) Connect(ctx context.Context) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	uri, err := url.Parse(db.cfg.Dsn)
	if err != nil {
		return err
	}

	var dialector gorm.Dialector
	if strings.HasPrefix(uri.Scheme, "sqlite") {
		dialector = sqlite.Open(uri.Host + uri.Path + uri.RawQuery)
	} else {
		dialector = postgres.Open(db.cfg.Dsn)
	}

	conn, err := gorm.Open(dialector, &gorm.Config{
		Logger: &SqlLogger{zap: xlogger.FromContext(ctx)},
	})
	if err != nil {
		return err
	}

	db.logger.Debug("connected")
	db.conn = conn
	return nil
}
