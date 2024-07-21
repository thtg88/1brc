package sequentialconsumer

import "github.com/thtg88/1brc/internal/models"

func (sc *SequentialConsumer) GetStats() map[string]models.CityStats {
	return sc.Stats
}
