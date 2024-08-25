package forecast

import (
	"encoding/json"
	"fmt"
	"go_weather_app/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetForecast(t *testing.T) {
	// Mock the /points API response
	pointsResponse := models.WeatherForecast{
		Properties: models.WeatherProperties{
			Forecast: "http://mocked-url.com/forecast",
		},
	}

	forecastResponse := models.ForecastResponse{
		Properties: models.ForecastProperties{
			Periods: []models.Period{
				{Temperature: 90, ShortForecast: "Sunny"},
				{Temperature: 85, ShortForecast: "Sunny"},
			},
		},
	}

	// Mock server for the /points API
	pointsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(pointsResponse)
	}))
	defer pointsServer.Close()

	// Mock server for the forecast URL
	forecastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the URL being requested is the correct one
		if r.URL.String() == "/forecast" {
			json.NewEncoder(w).Encode(forecastResponse)
		} else {
			http.Error(w, "not found", http.StatusNotFound)
		}
	}))
	defer forecastServer.Close()

	// Replace URLs in the test with our mock servers
	originalBaseURL := baseURL
	defer func() { baseURL = originalBaseURL }()
	baseURL = pointsServer.URL + "/"

	fmt.Println("Mock baseURL set to:", baseURL)

	// Overwrite the forecast URL in the mocked points response
	forecastResponseURL := forecastServer.URL + "/forecast"
	pointsResponse.Properties.Forecast = forecastResponseURL

	// Test the GetForecast function
	forecast, err := GetForecast(39.74, -104.99)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "The temperature is hot! Sunny"
	if forecast != expected {
		t.Errorf("Expected %s, got %s", expected, forecast)
	}
}
