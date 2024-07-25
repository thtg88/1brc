package bufferedsequentialconsumer

import "github.com/thtg88/1brc/internal/models"

func (bsc *BufferedSequentialConsumer) GetSortedStats() []models.CityStats {
	stats := bsc.GetStats()

	// Sort cities so that we are inserting stats in city's alphabetical order
	cities := bsc.GetSortedCities()

	sortedStats := make([]models.CityStats, len(cities))
	for idx, city := range cities {
		sortedStats[idx] = stats[city]
	}

	return sortedStats
}
