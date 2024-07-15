package repository

import "GolangAuthetication/internal/entity"

type IuserRepository interface {
	GetUser() ([]*entity.User, error)
	CreateUser(User *entity.User) error
	GetPerEmail(User *entity.User) (entity.User, error)
}
