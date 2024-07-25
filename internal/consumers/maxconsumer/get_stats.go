package maxconsumer

import "github.com/thtg88/1brc/internal/models"

func (mtc *MaxTempConsumer) GetStats() map[string]models.CityStats {
	stats := make(map[string]models.CityStats, len(mtc.Stats))
	for city, value := range mtc.Stats {
		stats[city] = models.CityStats{City: city, MaxTemp: value}
	}

	return stats
}
