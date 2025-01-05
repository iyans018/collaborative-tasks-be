package routes

import (
	"collaborative-task/handlers"
	"collaborative-task/repositories"
	"collaborative-task/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	api_v1 := app.Group("/api/v1")
	api_v1.Get("/users", handler.GetUsers)
	api_v1.Post("/users", handler.CreateUser)
}
