package sql

import (
	"gorm.io/gorm"
)

func UseWorkspace(ws string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("workspace_id = ?", ws)
	}
}
