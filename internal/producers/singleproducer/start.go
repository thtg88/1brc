package singleproducer

import (
	"io"
)

func (sp *SingleProducer) Start() {
	rowNumber := 1
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
			rowNumber++
			continue
		}
		if len(row) != 2 {
			sp.Logger.Println("row %d length not 2: %d", rowNumber, len(row))
			rowNumber++
			continue
		}
		// empty last row
		if row[1] == "" {
			rowNumber++
			continue
		}

		if sp.Config.Debug {
			sp.Logger.Printf("producing %v", row)
		}

		reading, err := sp.CSVRowParser.Parse(row)
		if err != nil {
			sp.Logger.Fatal(err)
		}

		sp.DataChannel <- reading

		sp.RecordsProduced++
		rowNumber++
	}

	sp.DoneChannel <- true
	close(sp.DataChannel)
	close(sp.DoneChannel)
}
