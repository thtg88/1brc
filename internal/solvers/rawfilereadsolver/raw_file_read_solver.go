package rawfilereadsolver

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/progressreporters"
	"github.com/thtg88/1brc/internal/resultswriters"
)

type RawFileReadSolver struct {
	Config           *configs.SolverConfig
	Logger           loggers.Logger
	ProgressReporter progressreporters.ProgressReporter
	ResultsWriter    resultswriters.Writer
}

func NewRawFileReadSolver(
	config *configs.SolverConfig,
	logger loggers.Logger,
	progressReporter progressreporters.ProgressReporter,
	resultsWriter resultswriters.Writer,
) *RawFileReadSolver {
	return &RawFileReadSolver{
		Config:           config,
		Logger:           logger,
		ProgressReporter: progressReporter,
		ResultsWriter:    resultsWriter,
	}
}
