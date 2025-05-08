package iRepo

import "github.com/x-sushant-x/Zocket/model"

type IUserRepository interface {
	CreateUser(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	GetAllUsers() []model.User
}
