// Package weather получение погоды по городу
package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather/geo"
)

// GetWeather получение погоды
func GetWeather(geo geo.Data, format int) string {
	baseURL, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseURL.RawQuery = params.Encode()
	resp, err := http.Get(baseURL.String())
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(body)
}
