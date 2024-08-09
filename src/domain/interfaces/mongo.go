package interfaces

import (
	"api.ddd/src/domain/aggregates"
)

type IRepository interface {
	CreateUser(example *aggregates.User) (*aggregates.User, error)
	UpdateUser(example *aggregates.User) (*aggregates.User, error)
	DeleteUser(exampleId string) error
}

type IFinder interface {
	GetUserById(exampleId string) (*aggregates.User, error)
	GetAllUsers() ([]*aggregates.User, error)
}
