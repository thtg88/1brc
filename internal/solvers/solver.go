package solvers

import (
	"os"

	"github.com/thtg88/1brc/internal/models"
)

type Solver interface {
	ProcessTemperatures(file *os.File) []models.CityStats
	Start()
	StopProgressReport()
}
