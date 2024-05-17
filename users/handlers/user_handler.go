package handlers

import (
	"user-services/entities"
	"user-services/services"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	var payload *entities.User = &entities.User{}
	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	result, err := h.userService.CreateUser(payload)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}
