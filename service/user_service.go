package service

import (
	"github.com/krifik/groupchat-be/model"
)

type UserService interface {
	Register(request model.CreateUserRequest) (response model.CreateUserResponse, err error)
	FindAll() ([]model.GetUserResponse, error)
	Login(request model.CreateUserRequest) (response model.CreateUserResponse, err error)
	GetUser(token string) (response model.GetUserResponse)
}
