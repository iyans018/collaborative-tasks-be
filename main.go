package main

import (
	"collaborative-task/config"
	"collaborative-task/db"
	"collaborative-task/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database
	db.Init()
	defer db.Close()

	// Initialize Fiber app
	app := fiber.New()

	// Register routes
	routes.RegisterUserRoutes(app)

	// Start the server
	log.Printf("Server is running on http://localhost:%s", config.AppPort)
	if err := app.Listen(":" + config.AppPort); err != nil {
		log.Fatal(err)
	}
}
