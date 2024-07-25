package avgconsumer

import "github.com/thtg88/1brc/internal/models"

func (atc *AvgTempConsumer) GetSortedStats() []models.CityStats {
	stats := atc.GetStats()

	// Sort cities so that we are inserting stats in city's alphabetical order
	cities := atc.GetSortedCities()

	sortedStats := make([]models.CityStats, len(cities))
	for idx, city := range cities {
		sortedStats[idx] = stats[city]
	}

	return sortedStats
}
