package sequentialconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
)

func TestSequentialConsumer_GetStats(t *testing.T) {
	t.Parallel()

	t.Run("no readings returns no stats", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildSequentialConsumer(&configs.SolverConfig{Debug: false}, logger)

		actualSortedStats := consumer.GetStats()

		require.Equal(t, map[string]models.CityStats{}, actualSortedStats)
	})

	t.Run("1 reading returns correct stats", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildSequentialConsumer(&configs.SolverConfig{Debug: false}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		consumer.ProcessReading(reading)

		expectedSortedStats := map[string]models.CityStats{
			builders.TemperatureReadingBuilder_TestCity: {
				City:              builders.TemperatureReadingBuilder_TestCity,
				MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
				MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
				AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
				MeasurementsSum:   builders.TemperatureReadingBuilder_TestTemperature,
				MeasurementsCount: 1,
			},
		}

		actualSortedStats := consumer.GetStats()

		require.Equal(t, expectedSortedStats, actualSortedStats)
	})

	t.Run("2 readings for the same city returns correct stats", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildSequentialConsumer(&configs.SolverConfig{Debug: false}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		consumer.ProcessReading(reading)
		consumer.ProcessReading(reading)

		expectedSortedStats := map[string]models.CityStats{
			builders.TemperatureReadingBuilder_TestCity: {
				City:              builders.TemperatureReadingBuilder_TestCity,
				MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
				MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
				AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
				MeasurementsSum:   2 * builders.TemperatureReadingBuilder_TestTemperature,
				MeasurementsCount: 2,
			},
		}

		actualSortedStats := consumer.GetStats()

		require.Equal(t, expectedSortedStats, actualSortedStats)
	})

	t.Run("2 readings for different cities returns correct stats", func(t *testing.T) {
		t.Parallel()

		const anotherCity = "another city"
		logger := loggermock.NewLoggerMock()
		consumer := buildSequentialConsumer(&configs.SolverConfig{Debug: false}, logger)
		reading1 := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		reading2 := builders.NewTemperatureReadingBuilder().WithCity(anotherCity).Build()
		consumer.ProcessReading(reading1)
		consumer.ProcessReading(reading2)

		expectedSortedStats := map[string]models.CityStats{
			builders.TemperatureReadingBuilder_TestCity: {
				City:              builders.TemperatureReadingBuilder_TestCity,
				MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
				MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
				AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
				MeasurementsSum:   builders.TemperatureReadingBuilder_TestTemperature,
				MeasurementsCount: 1,
			},
			anotherCity: {
				City:              anotherCity,
				MeasurementsCount: 1,
			},
		}

		actualSortedStats := consumer.GetStats()

		require.Equal(t, expectedSortedStats, actualSortedStats)
	})
}
