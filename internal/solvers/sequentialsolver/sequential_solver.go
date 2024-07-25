package sequentialsolver

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/progressreporters"
	"github.com/thtg88/1brc/internal/resultswriters"
)

type SequentialSolver struct {
	Config           *configs.SolverConfig
	ResultsWriter    resultswriters.Writer
	Logger           loggers.Logger
	ProgressReporter progressreporters.ProgressReporter
}

func NewSequentialSolver(config *configs.SolverConfig, resultsWriter resultswriters.Writer, logger loggers.Logger) *SequentialSolver {
	return &SequentialSolver{
		Config:           config,
		Logger:           logger,
		ProgressReporter: progressreporters.NewLogSleepReporter(logger, config.Progress),
		ResultsWriter:    resultsWriter,
	}
}
