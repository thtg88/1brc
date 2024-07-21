package csvrowprocessors

import "github.com/thtg88/1brc/internal/models"

type NilProcessor struct{}

func NewNilProcessor() *NilProcessor {
	return &NilProcessor{}
}

func (ncrp *NilProcessor) Process(row []string) (*models.TemperatureReading, error) {
	return nil, nil
}
