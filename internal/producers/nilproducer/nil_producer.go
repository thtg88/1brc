package nilproducer

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/csvreader"
	"github.com/thtg88/1brc/internal/csvrowparsers"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

type NilProducer struct {
	Config          *configs.SolverConfig
	CSVReader       csvreader.Reader
	CSVRowParser    csvrowparsers.Parser
	DataChannel     chan<- *models.TemperatureReading
	DoneChannel     chan<- bool
	Logger          loggers.Logger
	RecordsProduced uint64
}

func NewNilProducer(
	csvReader csvreader.Reader,
	dataChannel chan<- *models.TemperatureReading,
	doneChannel chan<- bool,
	logger loggers.Logger,
	config *configs.SolverConfig,
) *NilProducer {
	return &NilProducer{
		Config:       config,
		CSVReader:    csvReader,
		CSVRowParser: csvrowparsers.NewNilParser(),
		DataChannel:  dataChannel,
		DoneChannel:  doneChannel,
		Logger:       logger,
	}
}
