package services

import (
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
