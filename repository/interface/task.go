package iRepo

import "github.com/x-sushant-x/Zocket/model"

type ITaskRepository interface {
	CreateTask(task *model.Task) (*model.Task, error)
	GetAllTasks(status string) ([]model.Task, error)
	UpdateTaskStatus(taskID uint, newStatus string) error
	GetTaskByID(id uint) (*model.Task, error)
	AssignTask(taskID uint, userId uint) error
}
