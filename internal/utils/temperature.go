package utils

func CheckTemperature(temp int) string {
	if temp > 80 {
		return "The temperature is hot!"
	} else if temp < 60 {
		return "The temperature is cold!"
	} else {
		return "The temperature is moderate!"
	}
}
