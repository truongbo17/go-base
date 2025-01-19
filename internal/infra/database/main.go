package database

import "go-base/config"

func ConnectDatabase(configDatabase *config.DatabaseConnection) {
	ConnectMySQL(&configDatabase.DatabaseRelation)
	ConnectMongoDB(&configDatabase.DatabaseMongo)
}
