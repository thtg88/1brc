package maxconsumer

import (
	"fmt"

	"github.com/thtg88/1brc/internal/models"
)

func (mtc *MaxTempConsumer) GetCityStats(city string) (models.CityStats, error) {
	mtc.RLock()
	value, ok := mtc.Stats[city]
	mtc.RUnlock()
	if !ok {
		return models.CityStats{}, fmt.Errorf("%s city stats not found", city)
	}

	return models.CityStats{City: city, MaxTemp: value}, nil
}
