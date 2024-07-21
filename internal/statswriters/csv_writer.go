package statswriters

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/thtg88/1brc/internal/models"
)

type CSVWriter struct {
	Writer *csv.Writer
}

func NewCSVWriter(file io.Writer) Writer {
	return &CSVWriter{Writer: csv.NewWriter(file)}
}

func (csw *CSVWriter) Write(stats []models.CityStats) {
	csw.Writer.Write([]string{"city", "min", "max", "avg"})

	for _, cityStats := range stats {
		row := []string{
			cityStats.City,
			strconv.FormatFloat(float64(cityStats.MinTemp)/10, 'f', 1, 64),
			strconv.FormatFloat(float64(cityStats.MaxTemp)/10, 'f', 1, 64),
			strconv.FormatFloat(float64(cityStats.AverageTemp)/10, 'f', 1, 64),
		}

		csw.Writer.Write(row)
	}

	csw.Writer.Flush()
}
