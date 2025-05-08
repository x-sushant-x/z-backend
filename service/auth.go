package service

import (
	"errors"

	"github.com/x-sushant-x/Zocket/model"
	iRepo "github.com/x-sushant-x/Zocket/repository/interface"
	"github.com/x-sushant-x/Zocket/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo iRepo.IUserRepository
}

func NewAuthService(userRepo iRepo.IUserRepository) AuthService {
	return AuthService{
		userRepo: userRepo,
	}
}

func (s AuthService) Signup(name, email, password string) error {
	if name == "" || email == "" || password == "" {
		return errors.New("all fields are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.userRepo.CreateUser(user)
}

func (s AuthService) Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password required")
	}

	user, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
