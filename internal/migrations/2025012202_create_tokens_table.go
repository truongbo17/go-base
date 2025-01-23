package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateTokensTable = &gormigrate.Migration{
	ID: "2025012202",
	Migrate: func(tx *gorm.DB) error {
		return tx.Exec(`
			CREATE TABLE tokens (
				id INT AUTO_INCREMENT PRIMARY KEY,
				token VARCHAR(255) NOT NULL,
				type VARCHAR(20) NOT NULL,
				expires_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);
		`).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Exec("DROP TABLE IF EXISTS tokens;").Error
	},
}
