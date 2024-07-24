package csvreadermock

import "github.com/thtg88/1brc/internal/builders"

type CSVReaderMock struct{}

func NewCSVReaderMock() *CSVReaderMock {
	return &CSVReaderMock{}
}

func (crm *CSVReaderMock) Read() (record []string, err error) {
	return builders.NewTemperatureReadingCSVRowBuilder().WithTestValues().Build(), nil
}
