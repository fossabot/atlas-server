package client

import (
	"context"
	"sync"
	"time"

	"github.com/atlas-io/atlas-server/config"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var mongoOnce sync.Once
var mongoClient *mongo.Client

// GetMongoClient returns the current instance of the MongoDB client
func GetMongoClient() *mongo.Client {
	mongoOnce.Do(func() {
		connectToMongo()
	})

	return mongoClient
}

func GetMongoDefaultDatabase() *mongo.Database {
	mongoConfig := config.GetAppConfig().MongoConfig
	return GetMongoClient().Database(mongoConfig.Database)
}

func connectToMongo() {
	appConfig := config.GetAppConfig()
	mongoConfig := appConfig.MongoConfig
	ctx, cancelFn := context.WithTimeout(context.Background(), mongoConfig.ConnectTimeout*time.Second)
	mClient, err := mongo.Connect(ctx, mongoConfig.Server)
	if err != nil {
		cancelFn()
		panic(err)
	}
	mongoClient = mClient
	cancelFn()
}
