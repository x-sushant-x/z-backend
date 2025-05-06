package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string `json:"description"`
	Status      string `json:"status"`
	AssignedTo  uint   `json:"assignedTo"`
	User        User   `gorm:"foreignKey:AssignedTo" json:"user"`
}
