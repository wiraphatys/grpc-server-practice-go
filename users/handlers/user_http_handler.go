package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"user.services/entities"
	"user.services/services"
)

type userHttpHandler struct {
	userService services.UserService
}

func NewUserHttpHandler(userService services.UserService) UserHttpHandler {
	return &userHttpHandler{
		userService: userService,
	}
}

func (h *userHttpHandler) CreateUser(c *fiber.Ctx) error {
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

func (h *userHttpHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := strings.Trim(c.Params("email"), " ")

	result, err := h.userService.GetUserByEmail(email)
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
