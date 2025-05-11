package repository

import (
	"github.com/x-sushant-x/Zocket/model"
	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) CreateTask(task *model.Task) (*model.Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}

	var createdTask model.Task
	if err := r.db.Preload("User").First(&createdTask, task.ID).Error; err != nil {
		return nil, err
	}

	return &createdTask, nil
}

func (r *TaskRepo) GetAllTasks(status string) ([]model.Task, error) {
	var tasks []model.Task
	query := r.db.Preload("User")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query = query.Order("created_at DESC")

	err := query.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepo) UpdateTaskStatus(taskID uint, newStatus string) error {
	result := r.db.Model(&model.Task{}).Where("id = ?", taskID).Update("status", newStatus)
	return result.Error
}

func (r *TaskRepo) GetTaskByID(id uint) (*model.Task, error) {
	var task model.Task
	if err := r.db.Preload("User").First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
