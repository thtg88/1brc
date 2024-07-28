package sequentialsolver

import (
	"os"
	"time"
)

func (ss *SequentialSolver) Start() {
	file, err := os.Open(ss.Config.SourceFilePath)
	if err != nil {
		ss.Logger.Fatalf("could not open %s: %v", ss.Config.SourceFilePath, err)
	}
	defer file.Close()

	ss.Logger.Println("starting calculations...")

	start := time.Now()

	sortedStats, err := ss.ProcessTemperatures(file)
	if err != nil {
		ss.Logger.Fatalf("could not process temperatures: %v", err)
	}

	end := time.Now()

	duration := end.Sub(start)

	ss.Logger.Printf("calculations completed in %s!", duration.String())
	ss.Logger.Println("writing results...")

	err = ss.ResultsWriter.Write(sortedStats)
	if err != nil {
		ss.Logger.Fatalf("could not write results: %v", err)
	}
}
