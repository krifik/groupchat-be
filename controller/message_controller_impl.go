package controller

import (
	"mangojek-backend/exception"
	"mangojek-backend/helper"
	"mangojek-backend/model"
	"mangojek-backend/service"

	"github.com/gofiber/fiber/v2"
)

type MessageControllerImpl struct {
	MessageService service.MessageService
}

func NewMessageControllerImpl(messageService service.MessageService) MessageController {
	return &MessageControllerImpl{MessageService: messageService}
}

func (controller *MessageControllerImpl) SendMessage(c *fiber.Ctx) error {
	var request model.CreateMessageRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response := controller.MessageService.SendMessage(request)
	pusherClient := helper.NewPusherClient()
	pusherClient.Trigger("groupchat-channel", "message", response)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *MessageControllerImpl) GetMessages(c *fiber.Ctx) error {
	responses := controller.MessageService.GetMessages()
	if responses == nil {
		return c.Status(404).JSON(fiber.Map{
			"code":   404,
			"status": "NOT FOUND",
			"data":   []string{},
		})
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
