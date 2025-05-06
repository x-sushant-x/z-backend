package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/x-sushant-x/Zocket/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

func (con UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := con.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}
