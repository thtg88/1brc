package nilconsumer

import "github.com/thtg88/1brc/internal/models"

func (nc *NilConsumer) GetCityStats(_ string) (models.CityStats, error) {
	return models.CityStats{}, nil
}
