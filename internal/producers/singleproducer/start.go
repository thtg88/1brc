package singleproducer

import (
	"io"
)

func (sp *SingleProducer) Start() {
	for {
		if sp.RecordsProduced >= sp.Config.Limit {
			break
		}

		row, err := sp.CSVReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			sp.Logger.Fatalf("could not read from CSV file: %v", err)
		}
		if row == nil {
			continue
		}
		// empty last row
		if row[1] == "" {
			continue
		}

		if sp.Config.Debug {
			sp.Logger.Printf("producing %v", row)
		}

		reading, err := sp.CSVRowProcessor.Process(row)
		if err != nil {
			sp.Logger.Fatal(err)
		}

		sp.DataChannel <- reading

		sp.RecordsProduced++
	}

	sp.DoneChannel <- true
	close(sp.DataChannel)
	close(sp.DoneChannel)
}
