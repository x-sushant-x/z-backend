package service

import (
	"github.com/x-sushant-x/Zocket/model"
	"github.com/x-sushant-x/Zocket/repository"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (s UserService) GetAllUsers() ([]model.User, error) {
	users := s.userRepo.GetAllUsers()
	return users, nil
}
