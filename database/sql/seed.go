package sql

import (
	"bytes"
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/scrapnode/scrapcore/database"
	"gorm.io/gorm"
	"io"
	"os"
	"runtime"
	"strings"
)

func (db *SQL) Seed(ctx context.Context, seeds []string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if len(seeds) == 0 {
		db.Logger.Error(database.ErrSeedFilesEmpty.Error())
		return database.ErrSeedFilesEmpty
	}

	for _, filepath := range seeds {
		err := db.Conn.Transaction(func(tx *gorm.DB) error {
			// Force finalization of unreachable objects
			defer runtime.GC()

			file, err := os.Open(filepath)
			if err != nil {
				return err
			}

			b, err := io.ReadAll(file)
			if err != nil {
				return err
			}

			i := 0
			r := bytes.NewBuffer(b)
			for {
				stmt, err := r.ReadString(';')
				if err == io.EOF {
					// handle missing semicolon after last statement
					if strings.TrimSpace(stmt) != "" {
						err = nil
					} else {
						break
					}
				}
				if err != nil {
					return err
				}
				i++

				if txn := tx.Exec(stmt); txn.Error != nil {
					return txn.Error
				}
			}
			return nil
		})

		if err != nil {
			db.Logger.Error(database.ErrSeedFailed.Error())
			return err
		}
	}

	db.Logger.Debug("seeded")
	return nil
}
