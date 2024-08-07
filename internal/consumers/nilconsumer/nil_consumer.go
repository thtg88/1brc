package nilconsumer

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

type NilConsumer struct {
	Config          *configs.SolverConfig
	DataChannel     <-chan *models.TemperatureReading
	Logger          loggers.Logger
	RecordsConsumed uint64
	Stats           map[string]models.CityStats
}

func NewNilConsumer(
	dataChannel <-chan *models.TemperatureReading,
	logger loggers.Logger,
	config *configs.SolverConfig,
) *NilConsumer {
	return &NilConsumer{
		Config:      config,
		DataChannel: dataChannel,
		Logger:      logger,
		Stats:       make(map[string]models.CityStats),
	}
}
