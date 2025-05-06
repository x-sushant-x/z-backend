package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:100" json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}
