package minconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
)

func TestMinConsumer_GetStats(t *testing.T) {
	t.Parallel()

	t.Run("no readings returns no stats", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildMinConsumer(&configs.SolverConfig{}, logger)

		actualStats := consumer.GetStats()

		require.Equal(t, map[string]models.CityStats{}, actualStats)
	})

	t.Run("1 reading returns correct stats", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildMinConsumer(&configs.SolverConfig{}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		consumer.ProcessReading(reading)

		expectedStats := map[string]models.CityStats{
			builders.TemperatureReadingBuilder_TestCity: {
				City:    builders.TemperatureReadingBuilder_TestCity,
				MinTemp: builders.TemperatureReadingBuilder_TestTemperature,
			},
		}

		actualStats := consumer.GetStats()

		require.Equal(t, expectedStats, actualStats)
	})

	t.Run("2 readings for the same city returns correct stats", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildMinConsumer(&configs.SolverConfig{}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		consumer.ProcessReading(reading)
		consumer.ProcessReading(reading)

		expectedStats := map[string]models.CityStats{
			builders.TemperatureReadingBuilder_TestCity: {
				City:    builders.TemperatureReadingBuilder_TestCity,
				MinTemp: builders.TemperatureReadingBuilder_TestTemperature,
			},
		}

		actualStats := consumer.GetStats()

		require.Equal(t, expectedStats, actualStats)
	})

	t.Run("2 readings for different cities returns correct stats", func(t *testing.T) {
		t.Parallel()

		const anotherCity = "another city"
		logger := loggermock.NewLoggerMock()
		consumer := buildMinConsumer(&configs.SolverConfig{}, logger)
		reading1 := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		reading2 := builders.NewTemperatureReadingBuilder().WithCity(anotherCity).Build()
		consumer.ProcessReading(reading1)
		consumer.ProcessReading(reading2)

		expectedStats := map[string]models.CityStats{
			builders.TemperatureReadingBuilder_TestCity: {
				City:    builders.TemperatureReadingBuilder_TestCity,
				MinTemp: builders.TemperatureReadingBuilder_TestTemperature,
			},
			anotherCity: {
				City: anotherCity,
			},
		}

		actualStats := consumer.GetStats()

		require.Equal(t, expectedStats, actualStats)
	})
}
