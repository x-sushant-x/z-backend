package service

import (
	"github.com/x-sushant-x/Zocket/model"
	iRepo "github.com/x-sushant-x/Zocket/repository/interface"
)

type UserService struct {
	userRepo iRepo.IUserRepository
}

func NewUserService(userRepo iRepo.IUserRepository) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (s UserService) GetAllUsers() ([]model.User, error) {
	users := s.userRepo.GetAllUsers()
	return users, nil
}
