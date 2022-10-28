package controller

import (
	"mangojek-backend/exception"
	"mangojek-backend/helper"
	"mangojek-backend/model"
	"mangojek-backend/service"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Register(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response, err := controller.UserService.Register(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
func (controller *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	responses, err := controller.UserService.FindAll()
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *UserControllerImpl) Login(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response, err := controller.UserService.Login(request)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"code":   401,
			"status": "BAD REQUEST",
			"data":   "Wrong credential",
		})
	}

	claims := jwt.MapClaims{}
	claims["id"] = response.Id
	claims["name"] = response.Name
	claims["email"] = response.Email
	claims["image"] = response.Image
	claims["expired_at"] = time.Now().Add(72 * time.Hour).Unix()
	token, err := helper.GenerateJWT(&claims)
	exception.PanicIfNeeded(err)

	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}

func (controller *UserControllerImpl) GetUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	token = strings.Replace(token, "Bearer ", "", -1)
	user := controller.UserService.GetUser(token)
	return c.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	})
}
