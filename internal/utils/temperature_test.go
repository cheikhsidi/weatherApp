package utils

import "testing"

func TestCheckTemperature(t *testing.T) {
	tests := []struct {
		temp     int
		expected string
	}{
		{75, "It's hot outside!"},
		{65, "It's cold outside!"},
	}

	for _, test := range tests {
		result := CheckTemperature(test.temp)
		if result != test.expected {
			t.Errorf("CheckTemperature(%d) = %s; expected %s", test.temp, result, test.expected)
		}
	}
}
