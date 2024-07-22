package nilconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
)

func TestNilConsumer_ProcessReading(t *testing.T) {
	t.Parallel()

	t.Run("more than 0 readings does not add to stats", func(t *testing.T) {
		t.Parallel()

		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		mockLogger := loggermock.NewLoggerMock()
		consumer := buildNilConsumer(nil, mockLogger)
		wantStats := make(map[string]models.CityStats)

		consumer.ProcessReading(reading)
		actualStats := consumer.GetStats()

		require.Equal(t, wantStats, actualStats)
	})
}
