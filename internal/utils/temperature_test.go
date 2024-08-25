package utils

import "testing"

func TestCheckTemperature(t *testing.T) {
	tests := []struct {
		temp     int
		expected string
	}{
		{85, "It's hot outside!"},
		{70, "It's moderate outside!"},
		{55, "It's cold outside!"},
	}

	for _, test := range tests {
		if result := CheckTemperature(test.temp); result != test.expected {
			t.Errorf("CheckTemperature(%d) = %s, expected %s", test.temp, result, test.expected)
		}
	}
}
