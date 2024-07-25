package avgconsumer_test

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers/avgconsumer"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

func buildAvgConsumer(config *configs.SolverConfig, logger loggers.Logger) *avgconsumer.AvgTempConsumer {
	dataChannel := make(<-chan []*models.TemperatureReading)

	return avgconsumer.NewAvgTempConsumer(dataChannel, logger, config)
}
