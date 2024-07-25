package minconsumer

import (
	"fmt"

	"github.com/thtg88/1brc/internal/models"
)

func (mtc *MinTempConsumer) GetCityStats(city string) (models.CityStats, error) {
	minTemp, ok := mtc.Stats[city]
	if !ok {
		return models.CityStats{}, fmt.Errorf("%s city stats not found", city)
	}

	return models.CityStats{City: city, MinTemp: minTemp}, nil
}
