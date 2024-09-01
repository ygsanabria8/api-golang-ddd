package finder

import (
	"api.ddd/src/api/server"
	"api.ddd/src/domain/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type MongoUserFinder struct {
	database *mongo.Client
	logger   *server.Logger
	config   *server.Configuration
}

func NewMongoUserFinder(
	database *mongo.Client,
	logger *server.Logger,
	config *server.Configuration,
) interfaces.IUserFinder {
	return &MongoUserFinder{
		database: database,
		logger:   logger,
		config:   config,
	}
}

type MySqlUserFinder struct {
	database *gorm.DB
	logger   *server.Logger
}

func NewMySqlUserFinder(
	database *gorm.DB,
	logger *server.Logger,
) interfaces.IUserFinder {
	return &MySqlUserFinder{
		database: database,
		logger:   logger,
	}
}
