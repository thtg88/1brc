package temperatureswriter

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
)

type TemperaturesCSVWriter struct {
	csvWriter *csv.Writer
}

func NewTemperaturesCSVWriter(destinationFile *os.File) *TemperaturesCSVWriter {
	return &TemperaturesCSVWriter{
		csvWriter: csv.NewWriter(destinationFile),
	}
}

func (tcw *TemperaturesCSVWriter) Write(cities []string) {
	for i := 0; i < 1000000000; i++ {
		idx := rand.Intn(len(cities))
		city := cities[idx]
		// temp between -20 and 50 degrees Celsius
		temp := (rand.Float64() * 70) - 20
		row := []string{city, strconv.FormatFloat(temp, 'f', 1, 64)}

		tcw.csvWriter.Write(row)
	}
}
