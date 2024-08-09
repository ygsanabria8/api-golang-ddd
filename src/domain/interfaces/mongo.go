package interfaces

import (
	"api.ddd/src/domain/aggregates"
)

type IRepository interface {
	CreateUser(user *aggregates.User) (*aggregates.User, error)
	UpdateUser(user *aggregates.User) (*aggregates.User, error)
	DeleteUser(userId string) error
}

type IFinder interface {
	GetUserById(userId string) (*aggregates.User, error)
	GetAllUsers() ([]*aggregates.User, error)
}
