package main

import (
	"log"
	"os"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/solvers/bufferedsequentialsolver"
)

func main() {
	config := configs.NewDefaultSolverConfig()
	logger := log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	solver := bufferedsequentialsolver.NewBufferedSequentialSolver(config, logger)

	solver.Start()
}
