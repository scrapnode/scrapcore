package database

import "errors"

var (
	ErrQueryFailed    = errors.New("database: query is failed")
	ErrRecordNotFound = errors.New("database: record is not found")
)
