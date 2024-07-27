package sequentialsolver_test

import (
	"testing"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/mocks/resultwritermock"
	"github.com/thtg88/1brc/internal/progressreporters"
	"github.com/thtg88/1brc/internal/solvers/sequentialsolver"
)

func BenchmarkSequentialSolver_Start(b *testing.B) {
	config := configs.NewDefaultSolverConfig()
	config.Limit = uint64(b.N)
	logger := loggermock.NewLoggerMock()
	progressReporter := progressreporters.NewLogSleepReporter(logger, config.Progress)
	resultsWriter := resultwritermock.NewWriterMock()

	solver := sequentialsolver.NewSequentialSolver(config, logger, progressReporter, resultsWriter)

	solver.Start()
}
