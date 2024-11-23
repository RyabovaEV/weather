// Package geo геолакация
package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Data структура для получения погоды в указанном городе
type Data struct {
	City string `json:"city"`
}

// CityPopulationResponce структура для проверки города на валиность
type CityPopulationResponce struct {
	Error bool `json:"error"`
}

var ErrNoCity = errors.New("NOSITY")
var ErrNot200 = errors.New("NOT200")

// GetMyLocation получение погоды по локации или указанному городу
func GetMyLocation(city string) (*Data, error) {
	if city != "" {
		isCity := сheckCity(city)
		if !isCity {
			return nil, ErrNoCity
		}
		return &Data{
			City: city,
		}, nil
	}
	resp, err := http.Get("http://ipapi.co/json")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, ErrNot200
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo Data
	json.Unmarshal(body, &geo)
	return &geo, nil
}

// CheckCity проверка введенного города
func сheckCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var cpr CityPopulationResponce
	json.Unmarshal(body, &cpr)
	return !cpr.Error
}
