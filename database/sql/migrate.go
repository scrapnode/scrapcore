package sql

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/scrapnode/scrapcore/database"
)

func (db *SQL) Migrate(ctx context.Context) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.Configs.MigrateDir == "" {
		db.Logger.Error(database.ErrMigrationDirEmpty.Error())
		return database.ErrMigrationDirEmpty
	}

	dir := fmt.Sprintf("file://%s", db.Configs.MigrateDir)
	m, err := migrate.New(dir, db.Configs.Dsn)
	if err != nil {
		db.Logger.Errorw("could not construct migration", "directory", dir)
		return err
	}

	if err := m.Up(); err != nil {
		db.Logger.Errorw("migrate up got error", "error", err.Error())
		return err
	}
	return nil
}
