package bufferedsequentialconsumer

import (
	"fmt"

	"github.com/thtg88/1brc/internal/models"
)

func (bsc *BufferedSequentialConsumer) GetCityStats(city string) (models.CityStats, error) {
	cityStats, ok := bsc.Stats[city]
	if !ok {
		return models.CityStats{}, fmt.Errorf("%s city stats not found", city)
	}

	return models.CityStats{
		City:              city,
		MinTemp:           cityStats.MinTemp,
		MaxTemp:           cityStats.MaxTemp,
		AverageTemp:       cityStats.AverageTemp,
		MeasurementsSum:   cityStats.MeasurementsSum,
		MeasurementsCount: cityStats.MeasurementsCount,
	}, nil
}
