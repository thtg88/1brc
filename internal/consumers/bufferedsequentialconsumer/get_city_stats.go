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

	if !bsc.Config.CalculateAverageForEachReading {
		cityStats.AverageTemp = cityStats.MeasurementsSum / cityStats.MeasurementsCount
	}

	return cityStats, nil
}
