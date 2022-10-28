package repository

import (
	"github.com/krifik/groupchat-be/entity"
	"github.com/krifik/groupchat-be/model"
)

type MessageRepository interface {
	SendMessage(message *entity.Message)
	GetMessages() (createMessageResponse []model.CreateMessageResponse)
}
