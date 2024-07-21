package main

import (
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	sourceFile, err := os.Open("./data/weather_stations.csv")
	if err != nil {
		log.Fatalf("could not open weather_stations.csv: %v", err)
	}
	defer sourceFile.Close()

	log.Println("reading source cities file...")
	citiesMap := make(map[string]bool)
	csvReader := csv.NewReader(sourceFile)
	csvReader.Comma = ';'
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read from CSV file: %v", err)
		}
		if row == nil {
			continue
		}

		citiesMap[row[0]] = true
	}

	cities := make([]string, 0, len(citiesMap))
	for city := range citiesMap {
		cities = append(cities, city)
	}

	log.Println("cities file read!")
	log.Println("writing to temperature file...")

	destinationFile, err := os.Create("./data/temperatures.csv")
	if err != nil {
		log.Fatalf("could not create temperatures.csv: %v", err)
	}
	defer destinationFile.Close()

	csvWriter := csv.NewWriter(destinationFile)
	for i := 0; i < 1000000000; i++ {
		idx := rand.Intn(len(citiesMap))
		city := cities[idx]
		// temp between -20 and 50 degrees Celsius
		temp := (rand.Float64() * 70) - 20

		csvWriter.Write([]string{city, strconv.FormatFloat(temp, 'f', 1, 64)})
	}

	log.Println("temperatures file written!")
}
