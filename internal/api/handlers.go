package api

import (
	"go_weather_app/internal/forecast"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Weather App! Use the /forecast endpoint to get the weather forecast.")
	})

	// Define the GET /forecast endpoint
	app.Get("/forecast", handleForecast)
}

func handleForecast(c *fiber.Ctx) error {
	latitudeStr := c.Query("latitude")
	longitudeStr := c.Query("longitude")

	// Convert latitude and longitude from string to float64
	lat, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid latitude")
	}

	lon, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid longitude")
	}

	// Get the forecast data using the provided coordinates
	data, err := forecast.GetForecastFunc(lat, lon)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching forecast")
	}

	// Return the forecast data as the response
	return c.SendString(data)
}
