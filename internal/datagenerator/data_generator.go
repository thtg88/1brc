package datagenerator

import (
	"log"
	"os"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/temperatureswriter"
	"github.com/thtg88/1brc/internal/weatherstationsreader"
)

type DataGenerator struct {
	logger loggers.Logger
	config *configs.DataGeneratorConfig
}

func NewDataGenerator(logger loggers.Logger, config *configs.DataGeneratorConfig) *DataGenerator {
	return &DataGenerator{
		logger: logger,
		config: config,
	}
}

func (dg *DataGenerator) Generate() {
	sourceFile, err := os.Open(dg.config.SourceFilePath)
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

	destinationFile, err := os.Create(dg.config.DestinationFilePath)
	if err != nil {
		dg.logger.Fatalf("could not create temperatures.csv: %v", err)
	}
	defer destinationFile.Close()

	writer := temperatureswriter.NewTemperaturesCSVWriter(destinationFile)
	writer.Write(cities)

	dg.logger.Println("temperatures file written!")
}
