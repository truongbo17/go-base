package database

import (
	"context"
	"go-base/config"
	"go-base/internal/infra/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB(configDB *config.DatabaseMongo) *mongo.Client {
	if configDB.Uri != "" {
		logApp := logger.LogrusLogger

		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(configDB.Uri).SetServerAPIOptions(serverAPI)
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		MongoClient = client
		logApp.Infoln("Connect database mongodb success at " + configDB.Uri)
		if err != nil {
			panic(err)
		}

		return client
	}

	return nil
}
