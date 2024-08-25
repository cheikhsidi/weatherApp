package utils

func CheckTemperature(temp int) string {
	if temp > 70 {
		return "It's hot outside!"
	}
	return "It's cold outside!"
}
