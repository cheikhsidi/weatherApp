package forecast

import (
	"encoding/json"
	"fmt"
	"go_weather_app/internal/models"
	"go_weather_app/internal/utils"
	"log"
	"net/http"
)

var baseURL = "https://api.weather.gov/points/"

func GetForecast(latitude float64, longitude float64) (string, error) {
	// Construct the API URL with the given latitude and longitude
	url := fmt.Sprintf("%s%f,%f", baseURL, latitude, longitude)

	// Send an HTTP GET request to the API
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d", response.StatusCode)
		return "", fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	var weather models.WeatherForecast
	err = json.NewDecoder(response.Body).Decode(&weather)
	if err != nil {
		log.Println(err)
		return "", err
	}

	forecastURL := weather.Properties.Forecast

	response, err = http.Get(forecastURL)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer response.Body.Close()

	var forecast models.ForecastResponse
	err = json.NewDecoder(response.Body).Decode(&forecast)
	if err != nil {
		log.Println(err)
		return "", err
	}

	temp := forecast.Properties.Periods[1].Temperature
	message := utils.CheckTemperature(temp)

	return fmt.Sprintf("%s %s", message, forecast.Properties.Periods[1].ShortForecast), nil
}
