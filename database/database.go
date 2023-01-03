package database

import "context"

type Database interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Migrate(ctx context.Context) error
	Seed(ctx context.Context) error

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
