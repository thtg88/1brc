package rawfilereadproducer

import (
	"io"

	"github.com/thtg88/1brc/internal/models"
)

func (rfrp *RawFileReadProducer) Start() {
	carryover := ""
	readings := []*models.TemperatureReading{}
	for {
		if rfrp.RecordsProduced >= rfrp.Config.Limit {
			break
		}
		if rfrp.Config.FilePositioning.Enabled && rfrp.ReadUntilPosition >= rfrp.Config.FilePositioning.ReadUntilFilePosition {
			break
		}

		rows, newCarryover, err := rfrp.ReadRows(carryover)
		if err != nil {
			rfrp.Logger.Fatalf("could not read from CSV file: %v", err)
		}
		if len(rows) == 0 && len(newCarryover) == 0 {
			break
		}
		if rfrp.Config.FilePositioning.Enabled {
			currentOffset, err := rfrp.IOReadSeeker.Seek(0, io.SeekCurrent)
			if err != nil {
				rfrp.Logger.Fatalf("could get current offset: %v", err)
			}

			rfrp.ReadUntilPosition = currentOffset
		}

		for _, row := range rows {
			if rfrp.Config.Debug {
				rfrp.Logger.Printf("producing %v", row)
			}

			reading, err := rfrp.CSVRowParser.Parse(row)
			if err != nil {
				rfrp.Logger.Fatalf("could not process row %d: %v", rfrp.RecordsProduced+1, err)
			}

			readings = append(readings, reading)
			if len(readings) == int(rfrp.Config.BufferedChannelSize) {
				rfrp.DataChannel <- readings
				readings = []*models.TemperatureReading{}
			}

			rfrp.RecordsProduced++
		}

		carryover = newCarryover
	}
	// Flush last batch of readings
	if len(readings) > 0 {
		if rfrp.Config.Debug {
			rfrp.Logger.Println("flushing last batch of readings")
		}

		rfrp.DataChannel <- readings
	}

	close(rfrp.DataChannel)
}
