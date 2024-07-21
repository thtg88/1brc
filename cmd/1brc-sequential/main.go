package main

import (
	"log"
	"os"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/resultswriters"
	"github.com/thtg88/1brc/internal/solvers/sequentialsolver"
)

func main() {
	config := configs.NewDefaultSolverConfig()
	logger := log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	csvResultsWriter := resultswriters.NewCSVWriter(config.DestinationFilePath)
	solver := sequentialsolver.NewSequentialSolver(config, csvResultsWriter, logger)

	solver.Start()
}
