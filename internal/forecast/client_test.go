package forecast

import (
	"encoding/json"
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
				{Temperature: 75, ShortForecast: "Sunny"},
			},
		},
	}

	pointsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(pointsResponse)
	}))
	defer pointsServer.Close()

	forecastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(forecastResponse)
	}))
	defer forecastServer.Close()

	// Replace URLs in the test with our mock servers
	originalBaseURL := baseURL
	defer func() { baseURL = originalBaseURL }()
	baseURL = pointsServer.URL + "/"

	// Test the GetForecast function
	forecast, err := GetForecast(39.74, -104.99)
	if err != nil {
		// Log the error message and fail the test
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "It's hot outside! Sunny"
	if forecast != expected {
		t.Errorf("Expected %s, got %s", expected, forecast)
	}
}
