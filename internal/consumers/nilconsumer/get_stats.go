package nilconsumer

import "github.com/thtg88/1brc/internal/models"

func (nc *NilConsumer) GetStats() map[string]models.CityStats {
	return nc.Stats
}
