package maxconsumer_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
)

func TestMaxConsumer_GetCityStats(t *testing.T) {
	t.Parallel()

	t.Run("0 readings returns an error", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildMaxConsumer(&configs.SolverConfig{}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()

		expectedErr := fmt.Errorf("%s city stats not found", builders.TemperatureReadingBuilder_TestCity)

		actualCityStats, err := consumer.GetCityStats(reading.City)

		require.Equal(t, models.CityStats{}, actualCityStats)
		assert.Error(t, err)
		require.Equal(t, expectedErr, err)
	})

	t.Run("1 reading returns correct city stats", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildMaxConsumer(&configs.SolverConfig{}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		consumer.ProcessReading(reading)

		expectedCityStats := models.CityStats{
			City:    builders.TemperatureReadingBuilder_TestCity,
			MaxTemp: builders.TemperatureReadingBuilder_TestTemperature,
		}

		actualCityStats, err := consumer.GetCityStats(reading.City)

		require.NoError(t, err)
		require.Equal(t, expectedCityStats, actualCityStats)
	})

	t.Run("2 readings for the same city returns correct city stats", func(t *testing.T) {
		t.Parallel()

		logger := loggermock.NewLoggerMock()
		consumer := buildMaxConsumer(&configs.SolverConfig{}, logger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		consumer.ProcessReading(reading)
		consumer.ProcessReading(reading)

		expectedCityStats := models.CityStats{
			City:    builders.TemperatureReadingBuilder_TestCity,
			MaxTemp: builders.TemperatureReadingBuilder_TestTemperature,
		}

		actualCityStats, err := consumer.GetCityStats(reading.City)

		require.NoError(t, err)
		require.Equal(t, expectedCityStats, actualCityStats)
	})

	t.Run("2 readings for different cities returns correct city stats", func(t *testing.T) {
		t.Parallel()

		const anotherCity = "another city"
		logger := loggermock.NewLoggerMock()
		consumer := buildMaxConsumer(&configs.SolverConfig{}, logger)
		reading1 := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		reading2 := builders.NewTemperatureReadingBuilder().WithCity(anotherCity).Build()
		consumer.ProcessReading(reading1)
		consumer.ProcessReading(reading2)

		expectedCityStats := models.CityStats{
			City:    builders.TemperatureReadingBuilder_TestCity,
			MaxTemp: builders.TemperatureReadingBuilder_TestTemperature,
		}

		actualCityStats, err := consumer.GetCityStats(reading1.City)

		require.NoError(t, err)
		require.Equal(t, expectedCityStats, actualCityStats)
	})
}
