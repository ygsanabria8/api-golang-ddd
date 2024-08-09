package repository

import (
	"api.ddd/src/api/server"
	"api.ddd/src/domain/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	database *mongo.Client
	logger   *server.Logger
}

func NewMongoRepository(
	database *mongo.Client,
	logger *server.Logger,
) interfaces.IRepository {
	return &Repository{
		database: database,
		logger:   logger,
	}
}
