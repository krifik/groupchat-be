package repository

import (
	"mangojek-backend/config"
	"mangojek-backend/entity"
	"mangojek-backend/exception"
	"mangojek-backend/model"

	"gorm.io/gorm"
)

type MessageRepositoryImpl struct {
	DB *gorm.DB
}

func NewMessageRepositoryImpl(db *gorm.DB) MessageRepository {
	return &MessageRepositoryImpl{DB: db}
}

func (repository *MessageRepositoryImpl) GetMessages() (createMessageResponse []model.CreateMessageResponse) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	rows, err := repository.DB.WithContext(ctx).Omit("messages.*").Model(&entity.User{}).Select("messages.*, users.name, users.email, users.image").Joins("right join messages on messages.user_id = users.id").Rows()
	exception.PanicIfNeeded(err)
	defer rows.Close()
	for rows.Next() {
		result := model.CreateMessageResponse{}
		err := rows.Scan(&result.Id, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt, &result.UserID, &result.Content, &result.User.Name, &result.User.Email, &result.User.Image)
		exception.PanicIfNeeded(err)

		createMessageResponse = append(createMessageResponse, result)
	}
	return createMessageResponse
}
func (repository *MessageRepositoryImpl) SendMessage(message *entity.Message) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	result := repository.DB.WithContext(ctx).Create(&message)
	exception.PanicIfNeeded(result.Error)
}
