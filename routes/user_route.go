package routes

import (
	"collaborative-task/handlers"
	"collaborative-task/repositories"
	"collaborative-task/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterTaskRoutes(app *fiber.App) {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	app.Get("/users", handler.GetUsers)
	app.Post("/users", handler.CreateUser)
}
