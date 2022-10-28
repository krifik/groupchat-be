package service

import (
	"encoding/base64"
	"mangojek-backend/exception"
	"mangojek-backend/model"
	"mangojek-backend/repository"
	"mangojek-backend/validation"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) FindAll() ([]model.GetUserResponse, error) {
	users, _ := service.UserRepository.FindAll()
	var responses []model.GetUserResponse
	for _, user := range users {
		responses = append(responses, model.GetUserResponse{
			Id:        int(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
			Messages:  user.Messages,
		})
	}
	return responses, nil
}

func (service *UserServiceImpl) Register(request model.CreateUserRequest) (response model.CreateUserResponse, err error) {
	validation.Validate(request)
	result := service.UserRepository.CheckEmail(request)
	validation.IsEmailHasBeenTaken(result)
	decodedImage, err := base64.StdEncoding.DecodeString(request.Image[strings.IndexByte(request.Image, ',')+1:])
	exception.PanicIfNeeded(err)
	filename := strconv.Itoa(rand.Intn(100000)) + ".png"
	f, err := os.Create("app/storage/" + filename)
	exception.PanicIfNeeded(err)
	defer f.Close()
	request.Image = filename
	if _, err := f.Write(decodedImage); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	user, _ := service.UserRepository.Register(request)
	response = model.CreateUserResponse{
		Id:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Image:    user.Image,
	}
	return response, err
}

func (service *UserServiceImpl) Login(request model.CreateUserRequest) (response model.CreateUserResponse, err error) {
	validation.AuthValidate(request)
	user, err := service.UserRepository.Login(request)
	response = model.CreateUserResponse{
		Id:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Image:    user.Image,
	}

	return response, err
}

func (service *UserServiceImpl) GetUser(token string) (response model.GetUserResponse) {
	user := service.UserRepository.GetUser(token)
	response = model.GetUserResponse{
		Id:    int(user.Id),
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	}
	return response
}
