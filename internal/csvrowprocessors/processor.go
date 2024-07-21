package csvrowprocessors

import (
	"github.com/thtg88/1brc/internal/models"
)

type Processor interface {
	Process(row []string) (*models.TemperatureReading, error)
}
