package nilconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
)

func TestNilConsumer_GetCities(t *testing.T) {
	t.Parallel()

	t.Run("no readings returns no cities", func(t *testing.T) {
		t.Parallel()
		mockLogger := loggermock.NewLoggerMock()
		consumer := buildNilConsumer(nil, mockLogger)

		actualCities := consumer.GetCities()

		require.Equal(t, []string{}, actualCities)
	})

	t.Run("more than 0 readings returns no cities", func(t *testing.T) {
		t.Parallel()
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		mockLogger := loggermock.NewLoggerMock()
		consumer := buildNilConsumer(nil, mockLogger)
		consumer.ProcessReading(reading)

		actualCities := consumer.GetCities()

		require.Equal(t, []string{}, actualCities)
	})
}
