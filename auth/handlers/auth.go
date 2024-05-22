package handlers

import "github.com/gofiber/fiber/v2"

type AuthHttpHandler interface {
	Login(c *fiber.Ctx) error
}
