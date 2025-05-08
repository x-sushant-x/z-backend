package utils

import "github.com/gofiber/fiber/v2"

type ApiSuccess[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func SendApiSuccess[T any](c *fiber.Ctx, message string, data T) error {
	response := ApiSuccess[T]{
		Message: message,
		Data:    data,
	}
	return c.JSON(response)
}

type ApiError struct {
	Error string `json:"error"`
}

func SendApiError(c *fiber.Ctx, message string, status int) error {
	response := ApiError{
		Error: message,
	}
	return c.Status(status).JSON(response)
}
