package utils

import "testing"

func TestCheckTemperature(t *testing.T) {
	tests := []struct {
		temp     int
		expected string
	}{
		{85, "The temperature is hot!"},
		{70, "The temperature is moderate!"},
		{55, "The temperature is cold!"},
	}

	for _, test := range tests {
		if result := CheckTemperature(test.temp); result != test.expected {
			t.Errorf("CheckTemperature(%d) = %s, expected %s", test.temp, result, test.expected)
		}
	}
}
