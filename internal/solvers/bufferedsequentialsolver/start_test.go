package bufferedsequentialsolver_test

import (
	"testing"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers/bufferedsequentialconsumer"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
)

func BenchmarkSequentialSolver_Start(b *testing.B) {
	dataChannel := make(chan []*models.TemperatureReading)
	doneChannel := make(chan bool)
	config := configs.NewDefaultSolverConfig()
	config.Limit = uint64(b.N)
	logger := loggermock.NewLoggerMock()

	solver := bufferedsequentialconsumer.NewBufferedSequentialConsumer(dataChannel, doneChannel, logger, config)

	solver.Start()
}
