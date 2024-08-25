package forecast

import (
	"encoding/json"
	"fmt"
	"go_weather_app/internal/models"
	"go_weather_app/internal/utils"
	"io/ioutil"
	"net/http"
)

var baseURL = "https://api.weather.gov/points/"
var GetForecastFunc = GetForecast

func GetForecast(latitude float64, longitude float64) (string, error) {
	// Construct the API URL with the given latitude and longitude
	url := fmt.Sprintf("%s%f,%f", baseURL, latitude, longitude)
	fmt.Println("Requesting URL:", url)

	// Send an HTTP GET request to the API
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make request to the points API: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Error Body:", string(bodyBytes))
		return "", fmt.Errorf("points API request failed with status code: %d, body: %s", response.StatusCode, string(bodyBytes))
	}

	var weather models.WeatherForecast
	err = json.NewDecoder(response.Body).Decode(&weather)
	if err != nil {
		fmt.Println("Failed to decode points API response")
		return "", fmt.Errorf("failed to decode points API response: %w", err)
	}

	forecastURL := weather.Properties.Forecast
	fmt.Println("Forecast URL:", forecastURL)

	// Now, make the request to the forecast URL
	response, err = http.Get(forecastURL)
	if err != nil {
		return "", fmt.Errorf("failed to make request to the forecast API: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Error Body:", string(bodyBytes))
		return "", fmt.Errorf("forecast API request failed with status code: %d, body: %s", response.StatusCode, string(bodyBytes))
	}

	var forecast models.ForecastResponse
	err = json.NewDecoder(response.Body).Decode(&forecast)
	if err != nil {
		fmt.Println("Failed to decode forecast API response")
		return "", fmt.Errorf("failed to decode forecast API response: %w", err)
	}

	// Extract the desired forecast data
	temp := forecast.Properties.Periods[1].Temperature
	message := utils.CheckTemperature(temp)

	return fmt.Sprintf("%s %s", message, forecast.Properties.Periods[1].ShortForecast), nil
}
