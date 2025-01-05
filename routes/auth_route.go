package routes

import (
	"collaborative-task/handlers"
	"collaborative-task/repositories"
	"collaborative-task/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App) {
	repo := repositories.NewUserRepository()
	service := services.NewAuthService(repo)
	handler := handlers.NewAuthHandler(service)

	api_v1 := app.Group("/api/v1")
	api_v1.Post("/auth/register", handler.Register)
	api_v1.Post("/auth/login", handler.Login)
}
