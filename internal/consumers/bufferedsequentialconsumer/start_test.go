package bufferedsequentialconsumer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers/bufferedsequentialconsumer"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
)

func TestBufferedSequentialConsumer_Start(t *testing.T) {
	t.Parallel()

	type test struct {
		description             string
		expectedRecordsConsumed uint64
	}

	tests := []test{
		{
			description:             "no records consumed",
			expectedRecordsConsumed: 0,
		},
		{
			description:             "1 record consumed",
			expectedRecordsConsumed: 1,
		},
		{
			description:             "10 records consumed",
			expectedRecordsConsumed: 10,
		},
		{
			description:             "100 records consumed",
			expectedRecordsConsumed: 100,
		},
		{
			description:             "1000 records consumed",
			expectedRecordsConsumed: 1000,
		},
		{
			description:             "10000 records consumed",
			expectedRecordsConsumed: 10000,
		},
	}

	for _, testCase := range tests {
		tc := testCase
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			dataChannel := make(chan []*models.TemperatureReading)
			mockLogger := loggermock.NewLoggerMock()
			config := buildBufferedSequentialConsumerConfig(tc.expectedRecordsConsumed)

			consumer := bufferedsequentialconsumer.NewBufferedSequentialConsumer(dataChannel, mockLogger, config)

			go func() {
				readings := make([]*models.TemperatureReading, tc.expectedRecordsConsumed)
				for i := 0; i < int(tc.expectedRecordsConsumed); i++ {
					readings[i] = builders.NewTemperatureReadingBuilder().WithTestValues().Build()
				}

				dataChannel <- readings
				close(dataChannel)
			}()

			consumer.Start()

			actualRecordsConsumed := consumer.GetRecordsConsumed()

			require.Equal(t, tc.expectedRecordsConsumed, actualRecordsConsumed)
		})
	}
}

func buildBufferedSequentialConsumerConfig(limit uint64) *configs.SolverConfig {
	return &configs.SolverConfig{
		BufferedChannelSize: configs.DefaultBufferedChannelSize,
		Limit:               limit,
		Profile:             &configs.ProfileSolverConfig{},
		Progress:            &configs.ProgressSolverConfig{},
	}
}
