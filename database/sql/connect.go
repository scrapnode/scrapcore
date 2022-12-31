package sql

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

	dialector, err := Dial(db.Configs.Dsn)
	if err != nil {
		return err
	}

	conn, err := gorm.Open(dialector, &gorm.Config{
		Logger: &Logger{zap: xlogger.FromContext(ctx)},
	})
	if err != nil {
		return err
	}

	db.Logger.Debug("connected")
	db.Conn = conn
	return nil
}

// Dial support both SQLite & PostgreSQL
func Dial(dsn string) (gorm.Dialector, error) {
	uri, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(uri.Scheme, "sqlite") {
		return sqlite.Open(uri.Host + uri.Path + uri.RawQuery), nil
	}

	return postgres.Open(dsn), nil
}
