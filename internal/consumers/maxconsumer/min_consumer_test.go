package maxconsumer_test

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers/maxconsumer"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

func buildMaxConsumer(config *configs.SolverConfig, logger loggers.Logger) *maxconsumer.MaxTempConsumer {
	dataChannel := make(<-chan []*models.TemperatureReading)

	return maxconsumer.NewMaxTempConsumer(dataChannel, logger, config)
}
