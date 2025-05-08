package repository

import (
	"log"

	"github.com/x-sushant-x/Zocket/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

func (u UserRepo) CreateUser(user *model.User) error {
	return u.db.Create(user).Error
}

func (u UserRepo) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := u.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (u UserRepo) GetAllUsers() []model.User {
	var users []model.User

	result := u.db.Find(&users)
	if result.Error != nil {
		log.Printf("Error fetching users: %v", result.Error)
		return []model.User{}
	}

	return users
}
