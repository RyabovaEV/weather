// Package weather получение погоды по городу
package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather/geo"
)

type ErrResponce struct {
	Error bool `json:"error"`
}

var ErrPars = errors.New("ERR_URL")
var ErrGet = errors.New("ERR_HTTP")
var ErrRead = errors.New("ERR_READBODY")
var ErrFormat = errors.New("ERR_FORMAT")

// GetWeather получение погоды
func GetWeather(geo geo.Data, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrFormat
	}
	baseURL, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", ErrPars
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseURL.RawQuery = params.Encode()
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return "", ErrGet
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", ErrRead
	}
	return string(body), nil
}
