package service

import "github.com/krifik/groupchat-be/model"

type MessageService interface {
	SendMessage(request model.CreateMessageRequest) (response model.CreateMessageResponse)
	GetMessages() (responses []model.GetMessageResponse)
}
