package maxconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
)

func TestMaxConsumer_ProcessReading(t *testing.T) {
	t.Parallel()

	t.Run("successful reading processing", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildMaxConsumer(&configs.SolverConfig{Debug: true}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()

		expectedStats := map[string]int64{
			builders.TemperatureReadingBuilder_TestCity: builders.TemperatureReadingBuilder_TestTemperature,
		}

		consumer.ProcessReading(reading)
		consumer.ProcessReading(reading)
		actualStats := consumer.Stats
		actualRecordsConsumed := consumer.GetRecordsConsumed()

		require.Equal(t, expectedStats, actualStats)
		require.Equal(t, uint64(2), actualRecordsConsumed)
		require.Equal(t, uint64(2), logger.GetPrintfCalls())
	})

	t.Run("reading processing with debug off does not log", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildMaxConsumer(&configs.SolverConfig{}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()

		consumer.ProcessReading(reading)
		consumer.ProcessReading(reading)

		require.Zero(t, logger.GetPrintfCalls())
	})
}
