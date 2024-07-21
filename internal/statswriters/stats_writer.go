package statswriters

import "github.com/thtg88/1brc/internal/models"

type Writer interface {
	Write(stats []models.CityStats)
}
