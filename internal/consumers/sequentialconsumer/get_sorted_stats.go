package sequentialconsumer

import "github.com/thtg88/1brc/internal/models"

func (sc *SequentialConsumer) GetSortedStats() []models.CityStats {
	stats := sc.Stats

	// Sort cities so that we are inserting stats in city's alphabetical order
	cities := sc.GetSortedCities()

	sortedStats := make([]models.CityStats, len(cities))
	for idx, city := range cities {
		sortedStats[idx] = stats[city]
	}

	return sortedStats
}
