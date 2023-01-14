package database

import (
	"context"
)

func New(ctx context.Context, cfg *Configs) (Database, error) {
	// base con Dsn we will use different msgbus,
	// use SQL (SQLite/PostgreSQL) by default
	return NewSQL(ctx, cfg)
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
