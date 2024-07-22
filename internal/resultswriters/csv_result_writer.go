package resultswriters

import (
	"fmt"
	"os"

	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/statswriters"
)

type CSVWriter struct {
	DestinationFilePath string
}

func NewCSVWriter(destinationFilePath string) *CSVWriter {
	return &CSVWriter{DestinationFilePath: destinationFilePath}
}

func (cw *CSVWriter) Write(sortedStats []models.CityStats) error {
	destinationFile, err := os.Create(cw.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("could not create stats.csv: %v", err)
	}
	defer destinationFile.Close()

	statsWriter := statswriters.NewCSVWriter(destinationFile)
	statsWriter.Write(sortedStats)

	return nil
}
