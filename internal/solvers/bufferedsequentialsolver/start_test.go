package bufferedsequentialsolver_test

import (
	"testing"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/mocks/resultwritermock"
	"github.com/thtg88/1brc/internal/solvers/bufferedsequentialsolver"
)

func BenchmarkBufferedSequentialSolver_Start(b *testing.B) {
	config := configs.NewDefaultSolverConfig()
	config.Limit = uint64(b.N)
	logger := loggermock.NewLoggerMock()
	csvResultsWriter := resultwritermock.NewWriterMock()

	solver := bufferedsequentialsolver.NewBufferedSequentialSolver(config, logger, csvResultsWriter)

	solver.Start()
}
