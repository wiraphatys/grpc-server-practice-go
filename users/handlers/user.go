package handlers

import "github.com/gofiber/fiber/v2"

type UserHttpHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetUserByEmail(c *fiber.Ctx) error
}
