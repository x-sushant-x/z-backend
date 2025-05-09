package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/x-sushant-x/Zocket/requests"
	"github.com/x-sushant-x/Zocket/service"
	"github.com/x-sushant-x/Zocket/utils"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (con AuthController) Signup(c *fiber.Ctx) error {
	var req requests.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendApiError(c, "Invalid Input", fiber.StatusBadRequest)
	}

	err := req.Validate(true)
	if err != nil {
		return utils.SendApiError(c, err.Error(), fiber.StatusBadRequest)

	}

	err = con.authService.Signup(req.Name, req.Email, req.Password)
	if err != nil {
		return utils.SendApiError(c, err.Error(), fiber.StatusBadRequest)

	}

	return utils.SendApiSuccess(c, "Success", "Account Created")
}

func (con AuthController) Login(c *fiber.Ctx) error {
	var req requests.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendApiError(c, "Invalid Input", fiber.StatusBadRequest)
	}

	err := req.Validate(false)
	if err != nil {
		return utils.SendApiError(c, err.Error(), fiber.StatusBadRequest)

	}

	token, err := con.authService.Login(req.Email, req.Password)
	if err != nil {
		return utils.SendApiError(c, err.Error(), fiber.StatusBadRequest)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	})

	return utils.SendApiSuccess(c, "Success", token)
}
