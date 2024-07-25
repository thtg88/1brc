package main

import (
	"log"
	"os"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/progressreporters"
	"github.com/thtg88/1brc/internal/resultswriters"
	"github.com/thtg88/1brc/internal/solvers/bufferedsequentialsolver"
)

func main() {
	config := configs.NewDefaultSolverConfig()
	logger := log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	resultsWriter := resultswriters.NewCSVWriter(config.DestinationFilePath)
	progressReporter := progressreporters.NewLogSleepReporter(logger, config.Progress)
	solver := bufferedsequentialsolver.NewBufferedSequentialSolver(config, logger, progressReporter, resultsWriter)

	solver.Start()
}
