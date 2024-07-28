package bufferedproducer

import (
	"io"

	"github.com/thtg88/1brc/internal/models"
)

func (mp *BufferedProducer) Start() {
	readings := make([]*models.TemperatureReading, mp.Config.BufferedChannelSize)
	idx := 0
	rowNumber := 1
	for {
		if mp.RecordsProduced >= mp.Config.Limit {
			break
		}

		row, err := mp.CSVReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			mp.Logger.Fatalf("could not read from CSV file: %v", err)
		}
		if row == nil {
			rowNumber++
			continue
		}
		if len(row) != 2 {
			mp.Logger.Println("row %d length not 2: %d", rowNumber, len(row))
			rowNumber++
			continue
		}
		// empty last row
		if row[1] == "" {
			rowNumber++
			continue
		}

		if mp.Config.Debug {
			mp.Logger.Printf("producing %v", row)
		}

		reading, err := mp.CSVRowParser.Parse(row)
		if err != nil {
			mp.Logger.Fatal(err)
		}

		readings[idx] = reading

		mp.RecordsProduced++
		idx++

		if idx == int(mp.Config.BufferedChannelSize) {
			mp.DataChannel <- readings
			readings = make([]*models.TemperatureReading, mp.Config.BufferedChannelSize)
			idx = 0
		}

		rowNumber++
	}
	// Flush last batch of readings
	if idx > 0 {
		mp.DataChannel <- readings[0 : idx-1]
	}

	close(mp.DataChannel)
}
