package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/x-sushant-x/Zocket/service"
	"github.com/x-sushant-x/Zocket/utils"
)

type AuthRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (con AuthController) Signup(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendApiError(c, "Invalid Input", fiber.StatusBadRequest)
	}

	err := con.authService.Signup(req.Name, req.Email, req.Password)
	if err != nil {
		return utils.SendApiError(c, err.Error(), fiber.StatusBadRequest)

	}

	return utils.SendApiSuccess(c, "Success", "Account Created")
}

func (con AuthController) Login(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendApiError(c, "Invalid Input", fiber.StatusBadRequest)
	}

	token, err := con.authService.Login(req.Email, req.Password)
	if err != nil {
		return utils.SendApiError(c, err.Error(), fiber.StatusBadRequest)
	}

	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: token,
	})

	return utils.SendApiSuccess(c, "Success", token)
}
