package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	FindAll(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	// Insert(c *fiber.Ctx) error
}
