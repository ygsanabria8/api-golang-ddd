package repository

import (
	"api.ddd/src/api/server"
	"api.ddd/src/domain/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type MongoUserRepository struct {
	database *mongo.Client
	logger   *server.Logger
	config   *server.Configuration
}

func NewMongoUserRepository(
	database *mongo.Client,
	logger *server.Logger,
	config *server.Configuration,
) interfaces.IUserRepository {
	return &MongoUserRepository{
		database: database,
		logger:   logger,
		config:   config,
	}
}

type MySqlUserRepository struct {
	database *gorm.DB
	logger   *server.Logger
}

func NewUserMySqlRepository(
	database *gorm.DB,
	logger *server.Logger,
) interfaces.IUserRepository {
	return &MySqlUserRepository{
		database: database,
		logger:   logger,
	}
}
