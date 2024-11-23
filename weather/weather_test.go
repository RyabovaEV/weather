package weather_test

import (
	"strings"
	"testing"
	"weather/geo"
	"weather/weather"
)

func TestGetWeather(t *testing.T) {
	format := 3
	expected := "London"
	geoData := geo.Data{
		City: expected,
	}
	result, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}
	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "BigFormat", format: 147},
	{name: "0 format", format: 0},
	{name: "Minus Format", format: -1},
}

func TestGetWetherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "London"
			geoData := geo.Data{
				City: expected,
			}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrFormat {
				t.Errorf("Ожидалось %v, получено %v", weather.ErrFormat, err)
			}
		})
	}

}
