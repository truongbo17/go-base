package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateUsersTable = &gormigrate.Migration{
	ID: "2025012201",
	Migrate: func(tx *gorm.DB) error {
		return tx.Exec(`
			CREATE TABLE users (
				id INT AUTO_INCREMENT PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				email VARCHAR(100) UNIQUE NOT NULL,
				password VARCHAR(100) NULL DEFAULT '',
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);
		`).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Exec("DROP TABLE IF EXISTS users;").Error
	},
}
