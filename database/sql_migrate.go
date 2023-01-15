package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (db *SQL) Migrate(ctx context.Context) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.cfg.MigrateDir == "" {
		db.logger.Error(ErrMigrationDirNotSet.Error())
		return ErrMigrationDirNotSet
	}

	dir := fmt.Sprintf("file://%s", db.cfg.MigrateDir)
	m, err := migrate.New(dir, db.cfg.Dsn)
	if err != nil {
		db.logger.Errorw("could not construct migration", "directory", dir)
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		db.logger.Errorw("migrate up got error", "error", err.Error())
		return err
	}

	db.logger.Debug("migrated")
	return nil
}
