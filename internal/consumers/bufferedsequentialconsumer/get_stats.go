package bufferedsequentialconsumer

import "github.com/thtg88/1brc/internal/models"

func (bsc *BufferedSequentialConsumer) GetStats() map[string]models.CityStats {
	return bsc.Stats
}
