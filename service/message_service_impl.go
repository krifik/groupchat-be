package service

import (
	"mangojek-backend/entity"
	"mangojek-backend/model"
	"mangojek-backend/repository"
)

type MessageServiceImpl struct {
	MessageRepository repository.MessageRepository
}

func NewMessageRepositoryImpl(messageRepository repository.MessageRepository) MessageService {
	return &MessageServiceImpl{MessageRepository: messageRepository}
}

func (service *MessageServiceImpl) SendMessage(request model.CreateMessageRequest) (response model.CreateMessageResponse) {
	// validation here
	message := entity.Message{
		UserID:  request.UserID,
		Content: request.Content,
	}
	service.MessageRepository.SendMessage(&message)
	response = model.CreateMessageResponse{
		Id:        int(message.ID),
		UserID:    message.UserID,
		Content:   message.Content,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
		DeletedAt: message.DeletedAt,
		User:      request.User,
	}
	return response
}

func (service *MessageServiceImpl) GetMessages() (responses []model.GetMessageResponse) {
	createMessageResponse := service.MessageRepository.GetMessages()
	for _, message := range createMessageResponse {
		responses = append(responses, model.GetMessageResponse{
			Id:        int(message.Id),
			UserID:    message.UserID,
			Content:   message.Content,
			CreatedAt: message.CreatedAt,
			UpdatedAt: message.UpdatedAt,
			DeletedAt: message.DeletedAt,
			User:      message.User,
		})
	}
	return responses
}
