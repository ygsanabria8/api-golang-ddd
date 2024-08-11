package interfaces

import (
	"api.ddd/src/domain/aggregates"
)

type IUserRepository interface {
	CreateUser(user *aggregates.User) (*aggregates.User, error)
	UpdateUser(user *aggregates.User) (*aggregates.User, error)
	DeleteUser(userId string) error
}

type IUserFinder interface {
	GetUserById(userId string) (*aggregates.User, error)
	GetAllUsers() ([]*aggregates.User, error)
}

type IUserService interface {
	IUserFinder
	IUserRepository
}
