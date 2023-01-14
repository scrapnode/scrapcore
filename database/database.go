package database

import (
	"context"
	"github.com/scrapnode/scrapcore/database/configs"
	"github.com/scrapnode/scrapcore/database/sql"
)

func New(ctx context.Context, cfg *configs.Configs) (Database, error) {
	// base con Dsn we will use different msgbus,
	// use SQL (SQLite/PostgreSQL) by default
	return sql.New(ctx, cfg)
}

type Database interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Migrate(ctx context.Context) error
	Seed(ctx context.Context, seeds []string) error

	GetConn() any
}

type ListQuery struct {
	Cursor int
	Limit  int
}

type ListResult[T any] struct {
	Cursor  int
	Records []T
}
