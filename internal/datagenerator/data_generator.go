package datagenerator

import (
	"log"
	"os"

	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/temperatureswriter"
	"github.com/thtg88/1brc/internal/weatherstationsreader"
)

type DataGenerator struct {
	logger loggers.Logger
}

const (
	DestinationFilePath = "./data/temperatures.csv"
	SourceFilePath      = "./data/weather_stations.csv"
)

func NewDataGenerator(logger loggers.Logger) *DataGenerator {
	return &DataGenerator{logger: logger}
}

func (dg *DataGenerator) Generate() {
	sourceFile, err := os.Open(SourceFilePath)
	if err != nil {
		log.Fatalf("could not open weather_stations.csv: %v", err)
	}
	defer sourceFile.Close()

	dg.logger.Println("reading source cities file...")
	reader := weatherstationsreader.NewWeatherStationsCSVReader(sourceFile)
	cities, err := reader.Read()
	if err != nil {
		dg.logger.Fatal(err)
	}

	dg.logger.Println("cities file read!")
	dg.logger.Println("writing to temperature file...")

	destinationFile, err := os.Create(DestinationFilePath)
	if err != nil {
		dg.logger.Fatalf("could not create temperatures.csv: %v", err)
	}
	defer destinationFile.Close()

	writer := temperatureswriter.NewTemperaturesCSVWriter(destinationFile)
	writer.Write(cities)

	dg.logger.Println("temperatures file written!")
}
