package main

import (
	"go_weather_app/internal/forecast"
	"testing"
)

func TestMain(t *testing.T) {
	latitude := 39.74
	longitude := -104.99

	forecast, err := forecast.GetForecast(latitude, longitude)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Handle different possible outcomes
	expectedMessages := []string{
		"The temperature is hot! Sunny",
		"The temperature is hot! Partly Sunny then Slight Chance Showers And Thunderstorms",
	}

	found := false
	for _, expected := range expectedMessages {
		if forecast == expected {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Unexpected forecast: got %s", forecast)
	}
}
