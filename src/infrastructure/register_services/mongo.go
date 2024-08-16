package register_services

import (
	"api.ddd/src/api/server"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// ProvideMongoClient Provide Mongo Client Instance
func ProvideMongoClient(logger *server.Logger, config *server.Configuration) *mongo.Client {
	dsn := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		config.Mongo.User,
		config.Mongo.Password,
		config.Mongo.Host,
		config.Mongo.Port,
	)
	var clientOptions = options.Client().ApplyURI(dsn)

	clientOptions.SetMaxPoolSize(50)
	clientOptions.SetMinPoolSize(1)
	clientOptions.SetAppName(config.AppName)
	clientOptions.SetConnectTimeout(50 * time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logger.Fatalf("Cannot Connect With Mongo %v", err.Error())
		panic(err.Error())
	}

	logger.Debug("Connected To Mongo")
	return client
}
