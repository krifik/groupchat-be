package app

import (
	"github.com/krifik/groupchat-be/config"
	"github.com/krifik/groupchat-be/controller"
	"github.com/krifik/groupchat-be/repository"
	"github.com/krifik/groupchat-be/routes"
	"github.com/krifik/groupchat-be/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitializedApp() *fiber.App {
	configuration := config.NewConfiguration()
	database := config.NewPostgresDatabase(configuration)
	config.NewRunMigration(database)

	// Setup Repository
	userRepository := repository.NewUserRepositoryImpl(database)

	// Setup Service
	userService := service.NewUserServiceImpl(userRepository)

	// Setup Controller
	userController := controller.NewUserControllerImpl(userService)

	messageRepository := repository.NewMessageRepositoryImpl(database)
	messageService := service.NewMessageRepositoryImpl(messageRepository)
	messageController := controller.NewMessageControllerImpl(messageService)
	// Setup Fiber
	// app := fiber.New(config.NewFiberConfig())
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New(), cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Setup Routing
	routes.Route(app, userController, messageController)
	return app

}
