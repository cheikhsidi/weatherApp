package api

import (
	"go_weather_app/internal/forecast"
	"net/http"
	"net/http/httptest"
	"testing"

	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandleForecast(t *testing.T) {
	// Mock the forecast function to avoid real API calls during testing
	originalGetForecastFunc := forecast.GetForecastFunc
	forecast.GetForecastFunc = func(lat, lon float64) (string, error) {
		return "It's hot outside! Sunny", nil
	}
	defer func() {
		forecast.GetForecastFunc = originalGetForecastFunc
	}()

	// Create a new Fiber app instance
	app := fiber.New()
	RegisterRoutes(app)

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/forecast?latitude=39.74&longitude=-104.99", nil)

	// Record the HTTP response
	resp, err := app.Test(req, -1)

	// Assert that no error occurred
	assert.Nil(t, err)

	// Assert that the status code is 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	// Assert that the response body is as expected
	assert.Equal(t, "It's hot outside! Sunny", string(body))
}

func TestHandleForecastInvalidQuery(t *testing.T) {
	// Create a new Fiber app instance
	app := fiber.New()
	RegisterRoutes(app)

	// Create a new HTTP request with invalid latitude
	req := httptest.NewRequest("GET", "/forecast?latitude=invalid&longitude=-104.99", nil)

	// Record the HTTP response
	resp, err := app.Test(req, -1)

	// Assert that no error occurred
	assert.Nil(t, err)

	// Assert that the status code is 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestHandleForecastError(t *testing.T) {
	// Mock the forecast function to simulate an error
	originalGetForecastFunc := forecast.GetForecastFunc
	forecast.GetForecastFunc = func(lat, lon float64) (string, error) {
		return "", assert.AnError
	}
	defer func() {
		forecast.GetForecastFunc = originalGetForecastFunc
	}()

	// Create a new Fiber app instance
	app := fiber.New()
	RegisterRoutes(app)

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/forecast?latitude=39.74&longitude=-104.99", nil)

	// Record the HTTP response
	resp, err := app.Test(req, -1)

	// Assert that no error occurred
	assert.Nil(t, err)

	// Assert that the status code is 500 Internal Server Error
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
