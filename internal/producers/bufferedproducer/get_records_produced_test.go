package bufferedproducer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/mocks/csvreadermock"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/producers/bufferedproducer"
)

func TestBufferedProducer_GetRecordsProduced(t *testing.T) {
	t.Parallel()

	t.Run("new producer has produced 0 records", func(t *testing.T) {
		dataChannel := make(chan []*models.TemperatureReading, 1)
		producer := buildBufferedProducer(dataChannel, 0)

		actualRecordsProduced := producer.GetRecordsProduced()

		require.Zero(t, actualRecordsProduced)
	})

	t.Run("producer produced correct number records", func(t *testing.T) {
		type test struct {
			description             string
			expectedRecordsProduced uint64
		}

		tests := []test{
			{
				description:             "0 records produced",
				expectedRecordsProduced: 0,
			},
			{
				description:             "1 records produced",
				expectedRecordsProduced: 1,
			},
			{
				description:             "10 records produced",
				expectedRecordsProduced: 10,
			},
			{
				description:             "100 records produced",
				expectedRecordsProduced: 100,
			},
			{
				description:             "1000 records produced",
				expectedRecordsProduced: 1000,
			},
			{
				description:             "10000 records produced",
				expectedRecordsProduced: 10000,
			},
		}

		for _, testCase := range tests {
			tc := testCase
			t.Run(tc.description, func(t *testing.T) {
				dataChannel := make(chan []*models.TemperatureReading, 1)

				producer := buildBufferedProducer(dataChannel, tc.expectedRecordsProduced)

				go producer.Start()

				<-dataChannel

				actualRecordsProduced := producer.GetRecordsProduced()

				require.Equal(t, tc.expectedRecordsProduced, actualRecordsProduced)
			})
		}
	})
}

func buildBufferedProducer(dataChannel chan []*models.TemperatureReading, limit uint64) *bufferedproducer.BufferedProducer {
	csvReader := csvreadermock.NewCSVReaderMock()
	doneChannel := make(chan bool)
	mockLogger := loggermock.NewLoggerMock()
	config := &configs.SolverConfig{
		BufferedChannelSize: configs.DefaultBufferedChannelSize,
		Limit:               limit,
		Profile:             &configs.ProfileSolverConfig{},
		Progress:            &configs.ProgressSolverConfig{},
	}

	return bufferedproducer.NewBufferedProducer(csvReader, dataChannel, doneChannel, mockLogger, config)
}
