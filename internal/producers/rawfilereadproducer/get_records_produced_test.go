package rawfilereadproducer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
)

func TestRawFileReadProducer_GetRecordsProduced(t *testing.T) {
	t.Parallel()
	logger := loggermock.NewLoggerMock()
	producer := buildRawFileReadProducer(logger, 0, 0, false, false, false)

	actualRecordsProduced := producer.GetRecordsProduced()

	require.Zero(t, actualRecordsProduced)
}
