package geo_test

import (
	"testing"
	"weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "London"
	expected := geo.Data{
		City: city,
	}
	//Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	//Assert - проверка результата expected
	if err != nil {
		t.Error(err.Error())
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londonsss"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrNoCity, err)
	}
}
