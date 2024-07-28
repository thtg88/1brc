package sequentialconsumer_test

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers/sequentialconsumer"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

func buildSequentialConsumer(config *configs.SolverConfig, logger loggers.Logger) *sequentialconsumer.SequentialConsumer {
	dataChannel := make(<-chan *models.TemperatureReading)

	return sequentialconsumer.NewSequentialConsumer(dataChannel, logger, config)
}
