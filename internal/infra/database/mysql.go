package database

import (
	"fmt"
	"go-base/config"
	"go-base/internal/infra/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectMySQL(configDB *config.DatabaseRelation) *gorm.DB {
	if configDB.Username == "" {
		return nil
	}
	logApp := logger.LogrusLogger

	address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.Username,
		configDB.Password,
		configDB.Host,
		configDB.Port,
		configDB.Database,
	)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               address,
		DefaultStringSize: config.DefaultStringSizeMySql, // default size for string fields
	}), &gorm.Config{})

	if err != nil {
		panic("Connected failed, check your MySql")
	}

	// Migrate the schema
	//migrateErr := db.AutoMigrate(&models.Example{}, &models.User{})
	//if migrateErr != nil {
	//	panic(`Auto migrate failed, check your Mysql with ` + address)
	//}

	DB = db

	logApp.Infoln("Success connected to Mysql")

	return db
}
