package bufferedsequentialconsumer

import (
	"sync"

	"github.com/thtg88/1brc/internal/models"
)

func (bsc *BufferedSequentialConsumer) GetStats() map[string]models.CityStats {
	stats := bsc.Stats
	if bsc.Config.CalculateAverageForEachReading {
		return stats
	}

	var wg sync.WaitGroup
	var statsMutex sync.RWMutex
	statsMutex.RLock()
	for city, cityStats := range stats {
		wg.Add(1)
		go func(city string, cityStats models.CityStats) {
			defer func() {
				statsMutex.Unlock()
				wg.Done()
			}()

			cityStats.AverageTemp = cityStats.MeasurementsSum / cityStats.MeasurementsCount

			statsMutex.Lock()
			stats[city] = cityStats
		}(city, cityStats)
	}

	statsMutex.RUnlock()
	wg.Wait()

	return stats
}
