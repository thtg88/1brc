package sequentialconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
)

func TestSequentialConsumer_GetSortedStats(t *testing.T) {
	t.Parallel()

	t.Run("no readings returns no stats", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildSequentialConsumer(&configs.SolverConfig{}, logger)

		actualSortedStats := consumer.GetSortedStats()

		require.Equal(t, []models.CityStats{}, actualSortedStats)
	})

	t.Run("CalculateAverageForEachReading false", func(t *testing.T) {
		t.Run("1 reading returns correct stats", func(t *testing.T) {
			t.Parallel()

			logger := loggermock.NewLoggerMock()
			consumer := buildSequentialConsumer(&configs.SolverConfig{}, logger)
			reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
			consumer.ProcessReading(reading)

			expectedSortedStats := []models.CityStats{
				{
					City:              builders.TemperatureReadingBuilder_TestCity,
					MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsSum:   builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsCount: 1,
				},
			}

			actualSortedStats := consumer.GetSortedStats()

			require.Equal(t, expectedSortedStats, actualSortedStats)
		})

		t.Run("2 readings for the same city returns correct stats", func(t *testing.T) {
			t.Parallel()

			logger := loggermock.NewLoggerMock()
			consumer := buildSequentialConsumer(&configs.SolverConfig{}, logger)
			reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
			consumer.ProcessReading(reading)
			consumer.ProcessReading(reading)

			expectedSortedStats := []models.CityStats{
				{
					City:              builders.TemperatureReadingBuilder_TestCity,
					MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsSum:   2 * builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsCount: 2,
				},
			}

			actualSortedStats := consumer.GetSortedStats()

			require.Equal(t, expectedSortedStats, actualSortedStats)
		})

		t.Run("2 readings for different cities returns correct stats", func(t *testing.T) {
			t.Parallel()

			const anotherCity = "another city"
			logger := loggermock.NewLoggerMock()
			consumer := buildSequentialConsumer(&configs.SolverConfig{}, logger)
			reading1 := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
			reading2 := builders.NewTemperatureReadingBuilder().WithCity(anotherCity).Build()
			consumer.ProcessReading(reading1)
			consumer.ProcessReading(reading2)

			expectedSortedStats := []models.CityStats{
				{
					City:              anotherCity,
					MeasurementsCount: 1,
				},
				{
					City:              builders.TemperatureReadingBuilder_TestCity,
					MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsSum:   builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsCount: 1,
				},
			}

			actualSortedStats := consumer.GetSortedStats()

			require.Equal(t, expectedSortedStats, actualSortedStats)
		})
	})

	t.Run("CalculateAverageForEachReading true", func(t *testing.T) {
		t.Run("1 reading returns correct stats", func(t *testing.T) {
			t.Parallel()

			logger := loggermock.NewLoggerMock()
			consumer := buildSequentialConsumer(&configs.SolverConfig{CalculateAverageForEachReading: true}, logger)
			reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
			consumer.ProcessReading(reading)

			expectedSortedStats := []models.CityStats{
				{
					City:              builders.TemperatureReadingBuilder_TestCity,
					MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsSum:   builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsCount: 1,
				},
			}

			actualSortedStats := consumer.GetSortedStats()

			require.Equal(t, expectedSortedStats, actualSortedStats)
		})

		t.Run("2 readings for the same city returns correct stats", func(t *testing.T) {
			t.Parallel()

			logger := loggermock.NewLoggerMock()
			consumer := buildSequentialConsumer(&configs.SolverConfig{CalculateAverageForEachReading: true}, logger)
			reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
			consumer.ProcessReading(reading)
			consumer.ProcessReading(reading)

			expectedSortedStats := []models.CityStats{
				{
					City:              builders.TemperatureReadingBuilder_TestCity,
					MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsSum:   2 * builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsCount: 2,
				},
			}

			actualSortedStats := consumer.GetSortedStats()

			require.Equal(t, expectedSortedStats, actualSortedStats)
		})

		t.Run("2 readings for different cities returns correct stats", func(t *testing.T) {
			t.Parallel()

			const anotherCity = "another city"
			logger := loggermock.NewLoggerMock()
			consumer := buildSequentialConsumer(&configs.SolverConfig{CalculateAverageForEachReading: true}, logger)
			reading1 := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
			reading2 := builders.NewTemperatureReadingBuilder().WithCity(anotherCity).Build()
			consumer.ProcessReading(reading1)
			consumer.ProcessReading(reading2)

			expectedSortedStats := []models.CityStats{
				{
					City:              anotherCity,
					MeasurementsCount: 1,
				},
				{
					City:              builders.TemperatureReadingBuilder_TestCity,
					MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
					AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsSum:   builders.TemperatureReadingBuilder_TestTemperature,
					MeasurementsCount: 1,
				},
			}

			actualSortedStats := consumer.GetSortedStats()

			require.Equal(t, expectedSortedStats, actualSortedStats)
		})
	})
}
