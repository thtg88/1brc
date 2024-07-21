package consumers

import (
	"github.com/thtg88/1brc/internal/models"
)

type Consumer interface {
	GetCities() []string
	GetCityStats(city string) (models.CityStats, error)
	GetRecordsConsumed() uint64
	GetSortedCities() []string
	GetSortedStats() []models.CityStats
	GetStats() map[string]models.CityStats
	ProcessReading(reading *models.TemperatureReading)
	Start()
}
