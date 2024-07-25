package maxconsumer

import "github.com/thtg88/1brc/internal/models"

func (mtc *MaxTempConsumer) GetSortedStats() []models.CityStats {
	stats := mtc.Stats

	// Sort cities so that we are inserting stats in city's alphabetical order
	cities := mtc.GetSortedCities()

	sortedStats := make([]models.CityStats, len(cities))
	for idx, city := range cities {
		sortedStats[idx] = models.CityStats{City: city, MaxTemp: stats[city]}
	}

	return sortedStats
}
