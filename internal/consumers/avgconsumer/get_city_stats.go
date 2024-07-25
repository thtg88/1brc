package avgconsumer

import (
	"fmt"

	"github.com/thtg88/1brc/internal/models"
)

func (atc *AvgTempConsumer) GetCityStats(city string) (models.CityStats, error) {
	cityStats, ok := atc.Stats[city]
	if !ok {
		return models.CityStats{}, fmt.Errorf("%s city stats not found", city)
	}

	if !atc.Config.CalculateAverageForEachReading {
		cityStats.AverageTemp = cityStats.MeasurementsSum / cityStats.MeasurementsCount
	}

	return cityStats, nil
}
