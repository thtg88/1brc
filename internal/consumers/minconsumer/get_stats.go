package minconsumer

import "github.com/thtg88/1brc/internal/models"

func (mtc *MinTempConsumer) GetStats() map[string]models.CityStats {
	stats := make(map[string]models.CityStats, len(mtc.Stats))
	for city, minTemp := range mtc.Stats {
		stats[city] = models.CityStats{
			City:    city,
			MinTemp: minTemp,
		}
	}

	return stats
}
