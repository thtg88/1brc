package bufferedsequentialconsumer_test

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers/bufferedsequentialconsumer"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

func buildBufferedSequentialConsumer(config *configs.SolverConfig, logger loggers.Logger) *bufferedsequentialconsumer.BufferedSequentialConsumer {
	dataChannel := make(<-chan []*models.TemperatureReading)
	doneChannel := make(<-chan bool)

	return bufferedsequentialconsumer.NewBufferedSequentialConsumer(dataChannel, doneChannel, logger, config)
}
