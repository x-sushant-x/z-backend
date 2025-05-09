package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/x-sushant-x/Zocket/service"
	"github.com/x-sushant-x/Zocket/utils"
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
		return utils.SendApiError(c, "Failed to fetch users", fiber.StatusInternalServerError)

	}

	return utils.SendApiSuccess(c, "Users Fetched", users)
}
