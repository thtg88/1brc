package rawfilereadproducer_test

import (
	"fmt"
	"strings"

	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/mocks/ioreadseekermock"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/producers/rawfilereadproducer"
)

func buildRawFileReadProducer(logger loggers.Logger, bytesCount uint64, initialPosition uint64, debug bool, forceEOF bool, forceError bool) *rawfilereadproducer.RawFileReadProducer {
	rowsBytes := buildRowsBytes()
	ioReadSeeker := ioreadseekermock.NewIOReadSeekerMock(rowsBytes, initialPosition, forceEOF, forceError)
	dataChannel := make(chan []*models.TemperatureReading, 1)
	doneChannel := make(chan bool)
	config := &configs.SolverConfig{BufferedChannelSize: bytesCount, Debug: debug}

	return rawfilereadproducer.NewRawFileReadProducer(ioReadSeeker, dataChannel, doneChannel, logger, config)
}

func buildRowsBytes() []byte {
	rows := [][]string{
		builders.NewTemperatureReadingCSVRowBuilder().WithTestValues().Build(),
		builders.NewTemperatureReadingCSVRowBuilder().WithTestValues().Build(),
	}

	rowsStrings := make([]string, len(rows))
	for idx, row := range rows {
		rowsStrings[idx] = strings.Join(row, rawfilereadproducer.CommaSeparator)
	}

	rowsString := fmt.Sprintf(
		"%s%s",
		strings.Join(rowsStrings, rawfilereadproducer.NewLineSeparator),
		rawfilereadproducer.NewLineSeparator,
	)

	return []byte(rowsString)
}
