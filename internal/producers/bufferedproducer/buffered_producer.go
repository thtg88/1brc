package bufferedproducer

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/csvreader"
	"github.com/thtg88/1brc/internal/csvrowparsers"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

type BufferedProducer struct {
	Config          *configs.SolverConfig
	CSVReader       csvreader.Reader
	CSVRowParser    csvrowparsers.Parser
	DataChannel     chan<- []*models.TemperatureReading
	Logger          loggers.Logger
	RecordsProduced uint64
}

func NewBufferedProducer(
	csvReader csvreader.Reader,
	dataChannel chan<- []*models.TemperatureReading,
	logger loggers.Logger,
	config *configs.SolverConfig,
) *BufferedProducer {
	return &BufferedProducer{
		Config:       config,
		CSVReader:    csvReader,
		CSVRowParser: csvrowparsers.NewIntTempParser(),
		DataChannel:  dataChannel,
		Logger:       logger,
	}
}
