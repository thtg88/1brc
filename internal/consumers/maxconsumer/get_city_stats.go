package maxconsumer

import (
	"fmt"

	"github.com/thtg88/1brc/internal/models"
)

func (mtc *MaxTempConsumer) GetCityStats(city string) (models.CityStats, error) {
	maxTemp, ok := mtc.Stats[city]
	if !ok {
		return models.CityStats{}, fmt.Errorf("%s city stats not found", city)
	}

	return models.CityStats{City: city, MaxTemp: maxTemp}, nil
}
