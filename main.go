// Package main приложение для получения погоды в разных городах с помощью API
package main

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/weather"
)

func main() {

	fmt.Println("Погода в доме")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()
	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
