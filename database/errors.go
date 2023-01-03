package database

import "errors"

var (
	ErrQueryFailed        = errors.New("database: query is failed")
	ErrRecordNotFound     = errors.New("database: record is not found")
	ErrMigrationDirNotSet = errors.New("database: migration directory is not configured")
	ErrSeedFilesEmpty     = errors.New("database: seed files is empty")
	ErrSeedFailed         = errors.New("database: could not run seed file")
)
