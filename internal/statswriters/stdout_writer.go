package statswriters

import (
	"fmt"
	"io"
	"strconv"

	"github.com/thtg88/1brc/internal/models"
)

const CityStatsStdoutPattern = "%s:%s/%s/%s;"

type StdoutWriter struct {
	Writer io.Writer
}

func NewStdoutWriter(out io.Writer) Writer {
	return &StdoutWriter{Writer: out}
}

func (lsw *StdoutWriter) Write(stats []models.CityStats) {
	lsw.Writer.Write([]byte("{"))

	for _, cityStats := range stats {
		line := fmt.Sprintf(
			CityStatsStdoutPattern,
			cityStats.City,
			strconv.FormatFloat(float64(cityStats.MinTemp)/10, 'f', 1, 64),
			strconv.FormatFloat(float64(cityStats.MaxTemp)/10, 'f', 1, 64),
			strconv.FormatFloat(float64(cityStats.AverageTemp)/10, 'f', 1, 64),
		)

		lsw.Writer.Write([]byte(line))
	}

	lsw.Writer.Write([]byte("}"))
}
