package handlers

import (
	"context"

	"auth.services/entities"
	"auth.services/services"
	"github.com/gofiber/fiber/v2"
)

type authHttpHandler struct {
	authService services.AuthService
}

func NewAuthHttpHandler(authService services.AuthService) AuthHttpHandler {
	return &authHttpHandler{
		authService: authService,
	}
}

func (h *authHttpHandler) Login(c *fiber.Ctx) error {
	// create ctx
	ctx := context.Background()

	// parse request body into struct
	var reqBody *entities.LoginData
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	result, err := h.authService.VerifyUser(ctx, reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"result":  result,
	})
}
