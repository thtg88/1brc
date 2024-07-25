package sequentialconsumer

import (
	"fmt"

	"github.com/thtg88/1brc/internal/models"
)

func (sc *SequentialConsumer) GetCityStats(city string) (models.CityStats, error) {
	cityStats, ok := sc.Stats[city]
	if !ok {
		return models.CityStats{}, fmt.Errorf("%s city stats not found", city)
	}

	if !sc.Config.CalculateAverageForEachReading {
		cityStats.AverageTemp = cityStats.MeasurementsSum / cityStats.MeasurementsCount
	}

	return cityStats, nil
}
