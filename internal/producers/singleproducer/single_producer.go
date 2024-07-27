package singleproducer

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/csvreader"
	"github.com/thtg88/1brc/internal/csvrowprocessors"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

type SingleProducer struct {
	Config          *configs.SolverConfig
	CSVReader       csvreader.Reader
	CSVRowProcessor csvrowprocessors.Processor
	DataChannel     chan<- *models.TemperatureReading
	DoneChannel     chan<- bool
	Logger          loggers.Logger
	RecordsProduced uint64
}

func NewSingleProducer(
	csvReader csvreader.Reader,
	dataChannel chan<- *models.TemperatureReading,
	doneChannel chan<- bool,
	logger loggers.Logger,
	config *configs.SolverConfig,
) *SingleProducer {
	return &SingleProducer{
		Config:          config,
		CSVReader:       csvReader,
		CSVRowProcessor: csvrowprocessors.NewIntTempParseProcessor(),
		DataChannel:     dataChannel,
		DoneChannel:     doneChannel,
		Logger:          logger,
	}
}
