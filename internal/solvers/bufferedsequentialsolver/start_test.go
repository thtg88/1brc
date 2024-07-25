package bufferedsequentialsolver_test

import (
	"testing"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/mocks/resultwritermock"
	"github.com/thtg88/1brc/internal/progressreporters"
	"github.com/thtg88/1brc/internal/solvers/bufferedsequentialsolver"
)

func BenchmarkBufferedSequentialSolver_Start(b *testing.B) {
	config := configs.NewDefaultSolverConfig()
	config.Limit = uint64(b.N)
	logger := loggermock.NewLoggerMock()
	progressReporter := progressreporters.NewLogSleepReporter(logger, config.Progress)
	mockResultsWriter := resultwritermock.NewWriterMock()

	solver := bufferedsequentialsolver.NewBufferedSequentialSolver(config, logger, progressReporter, mockResultsWriter)

	solver.Start()
}
