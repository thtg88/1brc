package multirawfilereadsolver

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/progressreporters"
	"github.com/thtg88/1brc/internal/resultswriters"
)

type MultiRawFileReadSolver struct {
	Config           *configs.SolverConfig
	Logger           loggers.Logger
	ProgressReporter progressreporters.ProgressReporter
	ResultsWriter    resultswriters.Writer
}

func NewMultiRawFileReadSolver(
	config *configs.SolverConfig,
	logger loggers.Logger,
	progressReporter progressreporters.ProgressReporter,
	resultsWriter resultswriters.Writer,
) *MultiRawFileReadSolver {
	return &MultiRawFileReadSolver{
		Config:           config,
		Logger:           logger,
		ProgressReporter: progressReporter,
		ResultsWriter:    resultsWriter,
	}
}
