package mongo

import (
	"api.ddd/src/api/server"
	"api.ddd/src/domain/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type Finder struct {
	database *mongo.Client
	logger   *server.Logger
}

func NewMongoFinder(
	database *mongo.Client,
	logger *server.Logger,
) interfaces.IFinder {
	return &Finder{
		database: database,
		logger:   logger,
	}
}
