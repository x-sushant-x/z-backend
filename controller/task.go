package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/x-sushant-x/Zocket/service"
)

type TaskRequest struct {
	Description string `json:"description"`
	Status      string `json:"status"`
	AssignedTo  uint   `json:"assignedTo"`
}

type TaskController struct {
	taskService service.TaskService
}

type UpdateStatusRequest struct {
	TaskId uint   `json:"taskId"`
	Status string `json:"status"`
}

func NewTaskController(taskService service.TaskService) TaskController {
	return TaskController{taskService: taskService}
}

func (con TaskController) CreateTask(c *fiber.Ctx) error {
	var req TaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := con.taskService.CreateTask(req.Description, req.Status, req.AssignedTo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Task created"})
}

func (con TaskController) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := con.taskService.GetAllTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch tasks"})
	}

	return c.JSON(fiber.Map{"tasks": tasks})
}

func (con TaskController) UpdateTaskStatus(c *fiber.Ctx) error {
	var req UpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := con.taskService.UpdateTaskStatus(req.TaskId, req.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Task status updated successfully"})
}
