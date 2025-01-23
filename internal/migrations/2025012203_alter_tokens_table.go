package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AlterAddUserToTokensTable = &gormigrate.Migration{
	ID: "2025012203",
	Migrate: func(tx *gorm.DB) error {
		return tx.Exec(`
			ALTER TABLE tokens ADD COLUMN user INT DEFAULT 0;
		`).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
