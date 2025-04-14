package model

import "time"

// City представляет поддерживаемые города
type City string

const (
	CityMoscow City = "Москва"
	CitySPB    City = "Санкт-Петербург"
	CityKazan  City = "Казань"
)

// ValidCities содержит список поддерживаемых городов
var ValidCities = []City{CityMoscow, CitySPB, CityKazan}

// PickPoint представляет пункт выдачи заказов
type PickPoint struct {
	ID        int       `json:"id"`
	City      City      `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}

type PickPointFilter struct {
	StartDate *time.Time
	EndDate   *time.Time
}

type PickPointWithIntakes struct {
	PickPoint PickPoint
	Intakes   []Intake
}

// IsValidCity проверяет, является ли город поддерживаемым
func IsValidCity(city City) bool {
	for _, validCity := range ValidCities {
		if city == validCity {
			return true
		}
	}
	return false
}
