package controller

import "github.com/gofiber/fiber/v2"

type MessageController interface {
	GetMessages(c *fiber.Ctx) error
	SendMessage(c *fiber.Ctx) error
}
