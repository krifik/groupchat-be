package model

import (
	"mangojek-backend/entity"
	"time"

	"gorm.io/gorm"
)

type CreateUserRequest struct {
	Id int `json:"id"`
	// FirstName string `json:"first_name"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	Image     string           `json:"image"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at"`
	Messages  []entity.Message `json:"messages"`
}

type CreateUserResponse struct {
	Id int `json:"id"`
	// FirstName string `json:"first_name"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Password  string           `json:"-"`
	Image     string           `json:"image"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at"`
	Messages  []entity.Message `json:"messages"`
}

type GetUserResponse struct {
	Id int `json:"id"`
	// FirstName string `json:"first_name"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Password  string           `json:"-"`
	Image     string           `json:"image"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at"`
	Messages  []entity.Message `json:"messages"`
}
