package weatherstationsreader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type WeatherStationsCSVReader struct {
	csvReader *csv.Reader
}

func NewWeatherStationsCSVReader(sourceFile *os.File) *WeatherStationsCSVReader {
	csvReader := csv.NewReader(sourceFile)
	csvReader.Comma = ';'

	return &WeatherStationsCSVReader{csvReader: csvReader}
}

func (wscr *WeatherStationsCSVReader) Read() ([]string, error) {
	citiesMap := make(map[string]bool)

	for {
		row, err := wscr.csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("could not read from CSV file: %v", err)
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

	return cities, nil
}
