package sequentialconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
)

func TestSequentialConsumer_ProcessReading(t *testing.T) {
	t.Parallel()

	t.Run("successful reading processing", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildSequentialConsumer(&configs.SolverConfig{Debug: true}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()

		expectedStats := map[string]models.CityStats{
			builders.TemperatureReadingBuilder_TestCity: {
				City:              builders.TemperatureReadingBuilder_TestCity,
				MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
				MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
				AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
				MeasurementsSum:   2 * builders.TemperatureReadingBuilder_TestTemperature,
				MeasurementsCount: 2,
			},
		}

		consumer.ProcessReading(reading)
		consumer.ProcessReading(reading)
		actualStats := consumer.GetStats()
		actualRecordsConsumed := consumer.GetRecordsConsumed()

		require.Equal(t, expectedStats, actualStats)
		require.Equal(t, uint64(2), actualRecordsConsumed)
		require.Equal(t, uint64(2), logger.GetPrintfCalls())
	})

	t.Run("successful reading processing with debug off", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildSequentialConsumer(&configs.SolverConfig{}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()

		consumer.ProcessReading(reading)
		consumer.ProcessReading(reading)

		require.Equal(t, uint64(0), logger.GetPrintfCalls())
	})
}
