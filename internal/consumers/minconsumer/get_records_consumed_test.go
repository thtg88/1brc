package minconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
)

func TestMinConsumer_GetRecordsConsumed(t *testing.T) {
	t.Parallel()

	logger := loggermock.NewLoggerMock()
	consumer := buildMinConsumer(&configs.SolverConfig{}, logger)
	reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()

	consumer.ProcessReading(reading)
	actualRecordsConsumed := consumer.GetRecordsConsumed()

	require.Equal(t, uint64(1), actualRecordsConsumed)
}
