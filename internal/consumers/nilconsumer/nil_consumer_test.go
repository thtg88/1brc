package nilconsumer_test

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers/nilconsumer"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

func buildNilConsumer(config *configs.SolverConfig, logger loggers.Logger) *nilconsumer.NilConsumer {
	dataChannel := make(<-chan *models.TemperatureReading)
	doneChannel := make(<-chan bool)
	return nilconsumer.NewNilConsumer(dataChannel, doneChannel, logger, config)
}
