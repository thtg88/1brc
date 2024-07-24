package nilconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
)

func TestNilConsumer_GetRecordsConsumed(t *testing.T) {
	t.Parallel()
	t.Run("no readings returns 0 records consumed", func(t *testing.T) {
		t.Parallel()
		mockLogger := loggermock.NewLoggerMock()
		consumer := buildNilConsumer(nil, mockLogger)

		actualRecordsConsumed := consumer.GetRecordsConsumed()

		require.Zero(t, actualRecordsConsumed)
	})

	t.Run("n readings returns n records consumed", func(t *testing.T) {
		t.Parallel()
		mockLogger := loggermock.NewLoggerMock()
		consumer := buildNilConsumer(nil, mockLogger)
		reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
		consumer.ProcessReading(reading)

		actualRecordsConsumed := consumer.GetRecordsConsumed()

		require.Equal(t, uint64(1), actualRecordsConsumed)
	})
}
