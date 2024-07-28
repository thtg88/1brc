package singleproducer_test

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/mocks/csvreadermock"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/producers/singleproducer"
)

func buildSingleProducer(logger loggers.Logger, debug bool) *singleproducer.SingleProducer {
	csvReader := csvreadermock.NewCSVReaderMock()
	dataChannel := make(chan *models.TemperatureReading, 1)
	config := &configs.SolverConfig{Debug: debug}

	return singleproducer.NewSingleProducer(csvReader, dataChannel, logger, config)
}
