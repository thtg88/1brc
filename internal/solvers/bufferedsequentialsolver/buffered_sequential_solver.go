package bufferedsequentialsolver

import (
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/progressreporters"
	"github.com/thtg88/1brc/internal/resultswriters"
)

type BufferedSequentialSolver struct {
	Config           *configs.SolverConfig
	Logger           loggers.Logger
	ProgressReporter progressreporters.ProgressReporter
	ResultsWriter    resultswriters.Writer
}

func NewBufferedSequentialSolver(
	config *configs.SolverConfig,
	logger loggers.Logger,
	progressReporter progressreporters.ProgressReporter,
	resultsWriter resultswriters.Writer,
) *BufferedSequentialSolver {
	return &BufferedSequentialSolver{
		Config:           config,
		Logger:           logger,
		ProgressReporter: progressReporter,
		ResultsWriter:    resultsWriter,
	}
}
