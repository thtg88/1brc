package csvrowparsers

import "github.com/thtg88/1brc/internal/models"

type NilParser struct{}

func NewNilParser() *NilParser {
	return &NilParser{}
}

func (ncrp *NilParser) Parse(row []string) (*models.TemperatureReading, error) {
	return nil, nil
}
