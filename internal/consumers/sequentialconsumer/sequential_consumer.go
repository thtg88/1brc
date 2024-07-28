package sequentialconsumer

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

type SequentialConsumer struct {
	Config          *configs.SolverConfig
	DataChannel     <-chan *models.TemperatureReading
	Logger          loggers.Logger
	RecordsConsumed uint64
	Stats           map[string]models.CityStats
}

func NewSequentialConsumer(
	dataChannel <-chan *models.TemperatureReading,
	logger loggers.Logger,
	config *configs.SolverConfig,
) *SequentialConsumer {
	return &SequentialConsumer{
		Config:      config,
		DataChannel: dataChannel,
		Logger:      logger,
		Stats:       make(map[string]models.CityStats),
	}
}
