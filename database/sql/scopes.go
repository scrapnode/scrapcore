package sql

import (
	"github.com/benbjohnson/clock"
	"gorm.io/gorm"
)

func UseWorkspace(ws string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("workspace_id = ?", ws)
	}
}

func UseNotDeleted(c clock.Clock) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at = 0 OR deleted_at > ?", c.Now().UTC().UnixMilli())
	}
}
