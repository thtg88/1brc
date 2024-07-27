package rawfilereadproducer

import (
	"io"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/csvrowprocessors"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

const (
	CommaSeparator   = ","
	NewLineSeparator = "\n"
)

type RawFileReadProducer struct {
	Config          *configs.SolverConfig
	CSVRowProcessor csvrowprocessors.Processor
	IOReader        io.Reader
	DataChannel     chan<- []*models.TemperatureReading
	DoneChannel     chan<- bool
	Logger          loggers.Logger
	RecordsProduced uint64
}

func NewRawFileReadProducer(
	ioReader io.Reader,
	dataChannel chan<- []*models.TemperatureReading,
	doneChannel chan<- bool,
	logger loggers.Logger,
	config *configs.SolverConfig,
) *RawFileReadProducer {
	return &RawFileReadProducer{
		Config:          config,
		CSVRowProcessor: csvrowprocessors.NewIntTempParseProcessor(),
		IOReader:        ioReader,
		DataChannel:     dataChannel,
		DoneChannel:     doneChannel,
		Logger:          logger,
	}
}
