package resultswriters

import "github.com/thtg88/1brc/internal/models"

type Writer interface {
	Write(sortedStats []models.CityStats) error
}
