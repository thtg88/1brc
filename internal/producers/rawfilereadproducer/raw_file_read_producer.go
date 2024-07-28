package rawfilereadproducer

import (
	"io"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/csvrowparsers"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/models"
)

const (
	CommaSeparator   = ","
	NewLineSeparator = "\n"
)

type RawFileReadProducer struct {
	Config            *configs.SolverConfig
	CSVRowParser      csvrowparsers.Parser
	IOReadSeeker      io.ReadSeeker
	DataChannel       chan<- []*models.TemperatureReading
	DoneChannel       chan<- bool
	Logger            loggers.Logger
	ReadUntilPosition int64
	RecordsProduced   uint64
}

func NewRawFileReadProducer(
	ioReadSeeker io.ReadSeeker,
	dataChannel chan<- []*models.TemperatureReading,
	doneChannel chan<- bool,
	logger loggers.Logger,
	config *configs.SolverConfig,
) *RawFileReadProducer {
	return &RawFileReadProducer{
		Config:       config,
		CSVRowParser: csvrowparsers.NewIntTempParser(),
		IOReadSeeker: ioReadSeeker,
		DataChannel:  dataChannel,
		DoneChannel:  doneChannel,
		Logger:       logger,
	}
}
