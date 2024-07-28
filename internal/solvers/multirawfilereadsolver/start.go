package multirawfilereadsolver

import (
	"os"
	"time"
)

func (mrfrs *MultiRawFileReadSolver) Start() {
	file, err := os.Open(mrfrs.Config.SourceFilePath)
	if err != nil {
		mrfrs.Logger.Fatalf("could not open %s: %v", mrfrs.Config.SourceFilePath, err)
	}
	defer file.Close()

	mrfrs.Logger.Println("starting calculations...")

	start := time.Now()

	sortedStats, err := mrfrs.ProcessTemperatures(file)
	if err != nil {
		mrfrs.Logger.Fatalf("could not process temperatures: %v", err)
	}

	end := time.Now()

	duration := end.Sub(start)

	mrfrs.Logger.Printf("calculations completed in %s!", duration.String())
	mrfrs.Logger.Println("writing results...")

	err = mrfrs.ResultsWriter.Write(sortedStats)
	if err != nil {
		mrfrs.Logger.Fatalf("could not write results: %v", err)
	}
}
