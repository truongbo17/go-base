package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"go-base/config"
	"go-base/internal/infra/database"
	"go-base/internal/infra/logger"
)

func Migrate() {
	logger.Init()
	config.Init()
	configDatabase := config.EnvConfig.DatabaseConnection
	db := database.ConnectMySQL(&configDatabase.DatabaseRelation)

	migrationsList := []*gormigrate.Migration{
		CreateUsersTable,
		CreateTokensTable,
	}
	m := gormigrate.New(db, &gormigrate.Options{
		TableName:                 config.TableMigrate,
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            false,
		ValidateUnknownMigrations: false,
	}, migrationsList)
	if err := m.Migrate(); err != nil {
		panic(err)
	}

	logApp := logger.LogrusLogger
	logApp.Infoln("Migration successful.")
}
