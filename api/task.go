package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/x-sushant-x/Zocket/requests"
	"github.com/x-sushant-x/Zocket/service"
	"github.com/x-sushant-x/Zocket/utils"
)

type TaskController struct {
	taskService service.TaskService
}

func NewTaskController(taskService service.TaskService) TaskController {
	return TaskController{taskService: taskService}
}

func (con TaskController) CreateTask(c *fiber.Ctx) error {
	var req requests.TaskRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendApiError(c, "Invalid Request Body", fiber.StatusBadRequest)

	}

	err := con.taskService.CreateTask(req.Description, req.Status, req.AssignedTo)
	if err != nil {
		return utils.SendApiError(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SendApiSuccess(c, "Success", "Task Created")
}

func (con TaskController) GetAllTasks(c *fiber.Ctx) error {
	status := c.Query("status")

	tasks, err := con.taskService.GetAllTasks(status)
	if err != nil {
		return utils.SendApiError(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SendApiSuccess(c, "Success", tasks)

}

func (con TaskController) UpdateTaskStatus(c *fiber.Ctx) error {
	var req requests.UpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := con.taskService.UpdateTaskStatus(req.TaskId, req.Status)
	if err != nil {
		return utils.SendApiError(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SendApiSuccess(c, "Success", "Task Updated")
}

func (con TaskController) SuggestTasks(c *fiber.Ctx) error {
	suggestions, err := con.taskService.SuggestTasks()
	if err != nil {
		return utils.SendApiError(c, err.Error(), fiber.StatusInternalServerError)
	}

	return utils.SendApiSuccess(c, "Success", suggestions)
}
