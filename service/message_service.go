package service

import "mangojek-backend/model"

type MessageService interface {
	SendMessage(request model.CreateMessageRequest) (response model.CreateMessageResponse)
	GetMessages() (responses []model.GetMessageResponse)
}
