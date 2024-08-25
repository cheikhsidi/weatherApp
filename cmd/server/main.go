package main

import (
	"log"

	"go_weather_app/internal/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app instance
	app := fiber.New()

	// Register routes
	api.RegisterRoutes(app)

	// Start the Fiber app on port 3001
	log.Fatal(app.Listen(":3001"))
}
