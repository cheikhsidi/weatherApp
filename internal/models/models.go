package models

type WeatherForecast struct {
	Properties WeatherProperties `json:"properties"`
}

type WeatherProperties struct {
	Forecast       string `json:"forecast"`
	ForecastHourly string `json:"forecastHourly"`
}

type ForecastResponse struct {
	Properties ForecastProperties `json:"properties"`
}

type ForecastProperties struct {
	Periods []Period `json:"periods"`
}

type Period struct {
	Number           int    `json:"number"`
	Name             string `json:"name"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	IsDaytime        bool   `json:"isDaytime"`
	Temperature      int    `json:"temperature"`
	TemperatureUnit  string `json:"temperatureUnit"`
	WindSpeed        string `json:"windSpeed"`
	WindDirection    string `json:"windDirection"`
	Icon             string `json:"icon"`
	ShortForecast    string `json:"shortForecast"`
	DetailedForecast string `json:"detailedForecast"`
}
