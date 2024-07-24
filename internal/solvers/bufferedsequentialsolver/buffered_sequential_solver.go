package bufferedsequentialsolver

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/resultswriters"
)

type BufferedSequentialSolver struct {
	Config           *configs.SolverConfig
	CSVResultsWriter resultswriters.Writer
	Logger           loggers.Logger
}

func NewBufferedSequentialSolver(config *configs.SolverConfig, logger loggers.Logger) *BufferedSequentialSolver {
	return &BufferedSequentialSolver{
		Config:           config,
		CSVResultsWriter: resultswriters.NewCSVWriter(config.DestinationFilePath),
		Logger:           logger,
	}
}
