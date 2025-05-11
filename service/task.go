package service

import (
	"encoding/json"
	"errors"
	"github.com/x-sushant-x/Zocket/ai"

	"github.com/x-sushant-x/Zocket/constants"
	customErrors "github.com/x-sushant-x/Zocket/errors"
	"github.com/x-sushant-x/Zocket/model"
	iRepo "github.com/x-sushant-x/Zocket/repository/interface"
	"github.com/x-sushant-x/Zocket/socket"
)

type TaskService struct {
	taskRepo      iRepo.ITaskRepository
	wsClient      *socket.WebSocketClient
	userRepo      iRepo.IUserRepository
	aiSuggestions ai.Suggestions
}

func NewTaskService(taskRepo iRepo.ITaskRepository, wsClient *socket.WebSocketClient, userRepo iRepo.IUserRepository, aiSuggestions ai.Suggestions) TaskService {
	return TaskService{
		taskRepo:      taskRepo,
		wsClient:      wsClient,
		userRepo:      userRepo,
		aiSuggestions: aiSuggestions,
	}
}

func (s TaskService) CreateTask(description, status string, assignedTo *uint) error {
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

func (s TaskService) GetAllTasks(status string) ([]model.Task, error) {
	return s.taskRepo.GetAllTasks(status)
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

func (s TaskService) getTaskStats() (*model.TasksStats, error) {
	toDoTasks, err := s.taskRepo.GetAllTasks(constants.Task_ToDo)

	if err != nil {
		return nil, customErrors.ErrUnableToFetchTask
	}

	tasks := parseTasks(toDoTasks)

	users := s.userRepo.GetUsersWithStats()

	return &model.TasksStats{
		UsersStats:   users,
		NewTaskStats: tasks,
	}, nil
}

func parseTasks(toDoTasks []model.Task) []model.NewTasksStats {
	tasks := []model.NewTasksStats{}

	for _, task := range toDoTasks {
		taskStat := model.NewTasksStats{
			Title:          task.Description,
			EstimatedHours: task.EstimatedHours,
		}

		tasks = append(tasks, taskStat)
	}

	return tasks
}

func (s TaskService) SuggestTasks() ([]model.TaskAssignment, error) {
	stats, err := s.getTaskStats()
	if err != nil {
		return nil, customErrors.ErrInternalServerError
	}

	suggestions, err := s.aiSuggestions.SuggestTasks(stats)
	if err != nil {
		return nil, customErrors.ErrInternalServerError
	}

	return suggestions, nil
}
