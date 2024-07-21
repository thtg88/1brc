package nilproducer

import (
	"io"
)

func (np *NilProducer) Start() {
	for {
		if np.RecordsProduced >= np.Config.Limit {
			break
		}

		row, err := np.CSVReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			np.Logger.Fatalf("could not read from CSV file: %v", err)
		}
		if row == nil {
			continue
		}
		// empty last row
		if row[1] == "" {
			continue
		}

		if np.Config.Debug {
			np.Logger.Printf("producing %v", row)
		}

		_, err = np.CSVRowProcessor.Process(row)
		if err != nil {
			np.Logger.Fatal(err)
		}

		np.RecordsProduced++
	}
}
