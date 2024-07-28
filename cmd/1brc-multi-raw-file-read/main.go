package main

import (
	"log"
	"os"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/progressreporters"
	"github.com/thtg88/1brc/internal/resultswriters"
	"github.com/thtg88/1brc/internal/solvers/multirawfilereadsolver"
)

func main() {
	config := configs.NewDefaultSolverConfig()
	logger := log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	progressReporter := progressreporters.NewLogSleepReporter(logger, config.Progress)
	resultsWriter := resultswriters.NewCSVWriter(config.DestinationFilePath)
	solver := multirawfilereadsolver.NewMultiRawFileReadSolver(config, logger, progressReporter, resultsWriter)

	solver.Start()
}
