package bufferedsequentialconsumer

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

type BufferedSequentialConsumer struct {
	Config          *configs.SolverConfig
	DataChannel     <-chan []*models.TemperatureReading
	DoneChannel     <-chan bool
	Logger          loggers.Logger
	RecordsConsumed uint64
	Stats           map[string]models.CityStats
}

func NewBufferedSequentialConsumer(
	dataChannel <-chan []*models.TemperatureReading,
	doneChannel <-chan bool,
	logger loggers.Logger,
	config *configs.SolverConfig,
) *BufferedSequentialConsumer {
	return &BufferedSequentialConsumer{
		Config:      config,
		DataChannel: dataChannel,
		DoneChannel: doneChannel,
		Logger:      logger,
		Stats:       make(map[string]models.CityStats),
	}
}
