package repository

import (
	"fmt"
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
	existingUser, err := u.FindUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	if err := u.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (u UserRepo) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := u.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, result.Error
	}

	return &user, nil
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

func (u UserRepo) GetUsersWithStats() []model.UsersTaskStat {
	var stats []model.UsersTaskStat

	u.db.Model(&model.User{}).
		Select("users.name, COUNT(tasks.id) AS total_tasks_assigned, COALESCE(SUM(tasks.estimated_hours), 0) AS estimated_hours").
		Joins("LEFT JOIN tasks ON tasks.assigned_to = users.id").
		Group("users.id").
		Scan(&stats)

	return stats
}
