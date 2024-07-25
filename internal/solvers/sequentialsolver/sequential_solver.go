package sequentialsolver

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/progressreporters"
	"github.com/thtg88/1brc/internal/resultswriters"
)

type SequentialSolver struct {
	Config           *configs.SolverConfig
	Logger           loggers.Logger
	ProgressReporter progressreporters.ProgressReporter
	ResultsWriter    resultswriters.Writer
}

func NewSequentialSolver(
	config *configs.SolverConfig,
	logger loggers.Logger,
	progressReporter progressreporters.ProgressReporter,
	resultsWriter resultswriters.Writer,
) *SequentialSolver {
	return &SequentialSolver{
		Config:           config,
		Logger:           logger,
		ProgressReporter: progressReporter,
		ResultsWriter:    resultsWriter,
	}
}
