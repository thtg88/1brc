package csvrowprocessors

import (
	"fmt"

	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/parsers"
)

type IntTempParseProcessor struct{}

func NewIntTempParseProcessor() *IntTempParseProcessor {
	return &IntTempParseProcessor{}
}

func (itpp *IntTempParseProcessor) Process(row []string) (*models.TemperatureReading, error) {
	if len(row) != 2 {
		return nil, fmt.Errorf("row length not 2: %d", len(row))
	}

	temp, err := parsers.ParseTemperatureInt(row[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse temperature: %v", err)
	}

	city := row[0]

	reading := &models.TemperatureReading{
		City:        city,
		Temperature: temp,
	}

	return reading, nil
}
