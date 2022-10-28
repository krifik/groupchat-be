package repository

import (
	"mangojek-backend/entity"
	"mangojek-backend/model"
)

type MessageRepository interface {
	SendMessage(message *entity.Message)
	GetMessages() (createMessageResponse []model.CreateMessageResponse)
}
