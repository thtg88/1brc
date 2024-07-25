package maxconsumer

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

type MaxTempConsumer struct {
	Config          *configs.SolverConfig
	DataChannel     <-chan []*models.TemperatureReading
	Logger          loggers.Logger
	RecordsConsumed uint64
	Stats           map[string]int64
}

func NewMaxTempConsumer(
	dataChannel <-chan []*models.TemperatureReading,
	logger loggers.Logger,
	config *configs.SolverConfig,
) *MaxTempConsumer {
	return &MaxTempConsumer{
		Config:      config,
		DataChannel: dataChannel,
		Logger:      logger,
		Stats:       make(map[string]int64),
	}
}
