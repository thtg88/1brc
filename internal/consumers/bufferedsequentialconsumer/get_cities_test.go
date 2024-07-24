package bufferedsequentialconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
)

func TestBufferedSequentialConsumer_GetCities(t *testing.T) {
	t.Parallel()

	t.Run("0 readings returns 0 cities", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildBufferedSequentialConsumer(&configs.SolverConfig{}, logger)

		actualCities := consumer.GetCities()

		require.Equal(t, []string{}, actualCities)
	})

	t.Run("1 reading returns 1 city", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildBufferedSequentialConsumer(&configs.SolverConfig{}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		consumer.ProcessReading(reading)

		expectedCities := []string{builders.TemperatureReadingBuilder_TestCity}

		actualCities := consumer.GetCities()

		require.Equal(t, expectedCities, actualCities)
	})

	t.Run("2 readings for the same city returns 1 city", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildBufferedSequentialConsumer(&configs.SolverConfig{}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		consumer.ProcessReading(reading)
		consumer.ProcessReading(reading)

		expectedCities := []string{builders.TemperatureReadingBuilder_TestCity}

		actualCities := consumer.GetCities()

		require.Equal(t, expectedCities, actualCities)
	})

	t.Run("2 readings for different cities returns 2 cities", func(t *testing.T) {
		t.Parallel()

		const anotherCity = "another city"
		logger := loggermock.NewLoggerMock()
		consumer := buildBufferedSequentialConsumer(&configs.SolverConfig{}, logger)
		reading1 := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		reading2 := builders.NewTemperatureReadingBuilder().WithCity(anotherCity).Build()
		consumer.ProcessReading(reading1)
		consumer.ProcessReading(reading2)

		expectedCities := []string{
			builders.TemperatureReadingBuilder_TestCity,
			anotherCity,
		}

		actualCities := consumer.GetCities()

		require.Equal(t, expectedCities, actualCities)
	})
}
