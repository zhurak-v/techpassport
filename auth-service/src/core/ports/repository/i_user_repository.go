package repository

import (
	"github.com/google/uuid"
	"github.com/zhurak-v/techpassport/auth-service/src/core/entities"
)

type IUserRepositoryReader interface {
	FindAllUsers(withRelations bool, limit *int) (*[]entities.User, error)
	FindUserById(id uuid.UUID) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
}

type IUserRepositoryWriter interface {
	CreateUser(user *entities.User) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(user *entities.User) error
}

type IUserRepositoryContract interface {
	IUserRepositoryReader
	IUserRepositoryWriter
	BaseRepositoryContract
}
