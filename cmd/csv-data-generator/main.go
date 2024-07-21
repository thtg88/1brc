package main

import (
	"log"
	"os"

	"github.com/thtg88/1brc/internal/temperatureswriter"
	"github.com/thtg88/1brc/internal/weatherstationsreader"
)

func main() {
	sourceFile, err := os.Open("./data/weather_stations.csv")
	if err != nil {
		log.Fatalf("could not open weather_stations.csv: %v", err)
	}
	defer sourceFile.Close()

	log.Println("reading source cities file...")
	reader := weatherstationsreader.NewWeatherStationsCSVReader(sourceFile)
	cities, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("cities file read!")
	log.Println("writing to temperature file...")

	destinationFile, err := os.Create("./data/temperatures.csv")
	if err != nil {
		log.Fatalf("could not create temperatures.csv: %v", err)
	}
	defer destinationFile.Close()

	writer := temperatureswriter.NewTemperaturesCSVWriter(destinationFile)
	writer.Write(cities)

	log.Println("temperatures file written!")
}
