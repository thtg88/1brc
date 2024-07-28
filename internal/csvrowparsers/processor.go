package csvrowparsers

import (
	"github.com/thtg88/1brc/internal/models"
)

type Parser interface {
	Parse(row []string) (*models.TemperatureReading, error)
}
