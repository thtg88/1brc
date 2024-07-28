package csvrowparsers

import (
	"fmt"

	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/parsers"
)

type IntTempParser struct{}

func NewIntTempParser() *IntTempParser {
	return &IntTempParser{}
}

func (itpp *IntTempParser) Parse(row []string) (*models.TemperatureReading, error) {
	if len(row) != 2 {
		return nil, fmt.Errorf("row length not 2 (%d): %v", len(row), row)
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
