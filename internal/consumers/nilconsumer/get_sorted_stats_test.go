package nilconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
)

func TestNilConsumer_GetSortedStats(t *testing.T) {
	t.Parallel()

	t.Run("no readings returns no sorted stats", func(t *testing.T) {
		t.Parallel()
		mockLogger := loggermock.NewLoggerMock()
		consumer := buildNilConsumer(nil, mockLogger)

		actualSortedStats := consumer.GetSortedStats()

		require.Equal(t, []models.CityStats{}, actualSortedStats)
	})

	t.Run("more than 0 readings returns no sorted stats", func(t *testing.T) {
		t.Parallel()
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		mockLogger := loggermock.NewLoggerMock()
		consumer := buildNilConsumer(nil, mockLogger)
		consumer.ProcessReading(reading)

		actualSortedStats := consumer.GetSortedStats()

		require.Equal(t, []models.CityStats{}, actualSortedStats)
	})
}
