package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateTokensTable = &gormigrate.Migration{
	ID: "2025012201",
	Migrate: func(tx *gorm.DB) error {
		return tx.Exec(`
			CREATE TABLE tokens (
				id INT AUTO_INCREMENT PRIMARY KEY,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
			);
		`).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Exec("DROP TABLE IF EXISTS tokens;").Error
	},
}
