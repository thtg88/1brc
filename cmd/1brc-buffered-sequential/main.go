package main

import (
	"log"
	"os"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/resultswriters"
	"github.com/thtg88/1brc/internal/solvers/bufferedsequentialsolver"
)

func main() {
	config := configs.NewDefaultSolverConfig()
	logger := log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	csvResultsWriter := resultswriters.NewCSVWriter(config.DestinationFilePath)
	solver := bufferedsequentialsolver.NewBufferedSequentialSolver(config, logger, csvResultsWriter)

	solver.Start()
}
