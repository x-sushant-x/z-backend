package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description    string `json:"description"`
	Status         string `json:"status"`
	AssignedTo     uint   `json:"assignedTo"`
	User           User   `gorm:"foreignKey:AssignedTo" json:"user"`
	EstimatedHours int    `json:"estimatedHours"` // Days
}

type TasksStats struct {
	UsersStats   []UsersTaskStat `json:"users"`
	NewTaskStats []NewTasksStats `json:"tasks"`
}

type UsersTaskStat struct {
	Name               string `json:"name"`
	TotalTasksAssigned int    `json:"totalTasksAssigned"`
	EstimatedHours     int    `json:"estimatedHours"`
}

type NewTasksStats struct {
	Title          string `json:"title"`
	EstimatedHours int    `json:"estimatedHours"`
}

type TaskAssignment struct {
	User string `json:"user"`
	Task string `json:"task"`
}
