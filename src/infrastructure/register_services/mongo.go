package register_services

import (
	"api.ddd/src/api/server"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// ProvideMongoClient Provide Mongo Client Instance
func ProvideMongoClient(logger *server.Logger) *mongo.Client {
	var clientOptions = options.Client().ApplyURI("mongodb://admin:123456@localhost")
	clientOptions.SetMaxPoolSize(50)
	clientOptions.SetMinPoolSize(1)
	clientOptions.SetAppName("api-golang-ddd")
	clientOptions.SetConnectTimeout(50 * time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logger.Fatalf("Cannot Connect With Mongo %v", err.Error())
		panic(err.Error())
	}

	logger.Debug("Connected To Mongo")
	return client
}
