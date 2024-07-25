package avgconsumer

import "github.com/thtg88/1brc/internal/models"

func (atc *AvgTempConsumer) GetStats() map[string]models.CityStats {
	stats := atc.Stats
	if atc.Config.CalculateAverageForEachReading {
		return stats
	}

	fullStats := make(map[string]models.CityStats, len(stats))
	for city, cityStats := range atc.Stats {
		fullStats[city] = models.CityStats{
			City:              city,
			AverageTemp:       cityStats.MeasurementsSum / cityStats.MeasurementsCount,
			MeasurementsSum:   cityStats.MeasurementsSum,
			MeasurementsCount: cityStats.MeasurementsCount,
		}
	}

	return fullStats
}
