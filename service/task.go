package service

import (
	"encoding/json"
	"errors"

	"github.com/x-sushant-x/Zocket/model"
	iRepo "github.com/x-sushant-x/Zocket/repository/interface"
	"github.com/x-sushant-x/Zocket/socket"
)

type TaskService struct {
	taskRepo iRepo.ITaskRepository
	wsClient *socket.WebSocketClient
}

func NewTaskService(taskRepo iRepo.ITaskRepository, wsClient *socket.WebSocketClient) TaskService {
	return TaskService{
		taskRepo: taskRepo,
		wsClient: wsClient,
	}
}

func (s TaskService) CreateTask(description, status string, assignedTo uint) error {
	if description == "" || status == "" {
		return errors.New("all fields are required")
	}

	task := &model.Task{
		Description: description,
		Status:      status,
		AssignedTo:  assignedTo,
	}

	createdTask, err := s.taskRepo.CreateTask(task)

	if err != nil {
		return err
	}

	taskJSON, err := json.Marshal(createdTask)
	if err == nil {
		s.wsClient.Broadcast(taskJSON)
	}

	return nil
}

func (s TaskService) GetAllTasks() ([]model.Task, error) {
	return s.taskRepo.GetAllTasks()
}

func (s TaskService) UpdateTaskStatus(taskID uint, newStatus string) error {
	if newStatus == "" {
		return errors.New("status is required")
	}
	err := s.taskRepo.UpdateTaskStatus(taskID, newStatus)

	if err != nil {
		return err
	}

	task, err := s.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return err
	}

	taskJSON, err := json.Marshal(task)
	if err == nil {
		s.wsClient.Broadcast(taskJSON)
	}

	return nil
}
