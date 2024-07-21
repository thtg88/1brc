package nilconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
)

func TestNilConsumer_GetCityStats(t *testing.T) {
	t.Parallel()

	t.Run("no readings returns no city stats", func(t *testing.T) {
		t.Parallel()
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		mockLogger := loggermock.NewLoggerMock()
		consumer := buildNilConsumer(nil, mockLogger)

		actualCityStats, actualErr := consumer.GetCityStats(reading.City)

		require.Equal(t, models.CityStats{}, actualCityStats)
		require.Equal(t, nil, actualErr)
	})

	t.Run("more than 0 readings returns no city stats", func(t *testing.T) {
		t.Parallel()

		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		mockLogger := loggermock.NewLoggerMock()
		consumer := buildNilConsumer(nil, mockLogger)
		consumer.ProcessReading(reading)

		actualCityStats, actualErr := consumer.GetCityStats(reading.City)

		require.Equal(t, models.CityStats{}, actualCityStats)
		require.Equal(t, nil, actualErr)
	})
}
