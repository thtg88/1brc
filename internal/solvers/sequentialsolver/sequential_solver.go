package sequentialsolver

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/resultswriters"
)

type SequentialSolver struct {
	Config        *configs.SolverConfig
	ResultsWriter resultswriters.Writer
	Logger        loggers.Logger
}

func NewSequentialSolver(config *configs.SolverConfig, resultsWriter resultswriters.Writer, logger loggers.Logger) *SequentialSolver {
	return &SequentialSolver{
		Config:        config,
		ResultsWriter: resultsWriter,
		Logger:        logger,
	}
}
