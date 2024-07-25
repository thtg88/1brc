package minconsumer_test

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers/minconsumer"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

func buildMinConsumer(config *configs.SolverConfig, logger loggers.Logger) *minconsumer.MinTempConsumer {
	dataChannel := make(<-chan []*models.TemperatureReading)

	return minconsumer.NewMinTempConsumer(dataChannel, logger, config)
}
