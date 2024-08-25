package utils

func CheckTemperature(temp int) string {
	if temp > 80 {
		return "It's hot outside!"
	} else if temp < 60 {
		return "It's cold outside!"
	} else {
		return "It's moderate outside!"
	}
}
