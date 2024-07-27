package rawfilereadsolver

import (
	"os"
	"time"
)

func (rfrs *RawFileReadSolver) Start() {
	file, err := os.Open(rfrs.Config.SourceFilePath)
	if err != nil {
		rfrs.Logger.Fatalf("could not open %s: %v", rfrs.Config.SourceFilePath, err)
	}
	defer file.Close()

	rfrs.Logger.Println("starting calculations...")

	start := time.Now()

	sortedStats := rfrs.ProcessTemperatures(file)

	end := time.Now()

	duration := end.Sub(start)

	rfrs.Logger.Printf("calculations completed in %s!", duration.String())
	rfrs.Logger.Println("writing results...")

	err = rfrs.ResultsWriter.Write(sortedStats)
	if err != nil {
		rfrs.Logger.Fatalf("%v", err)
	}
}
