package progressreporters_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers/nilconsumer"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/mocks/csvreadermock"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/producers/nilproducer"
	"github.com/thtg88/1brc/internal/progressreporters"
)

func TestLogSleepReporter_ConsumerReport(t *testing.T) {
	t.Parallel()

	t.Run("ShouldReport config set to false does not log anything", func(t *testing.T) {
		t.Parallel()

		config := buildNilConsumerConfig(false)
		mockLogger := loggermock.NewLoggerMock()
		progressReporter := progressreporters.NewLogSleepReporter(mockLogger, config.Progress)
		consumer := buildNilConsumer(config, mockLogger)

		progressReporter.ConsumerReport(consumer)

		require.Equal(t, mockLogger.GetPrintfCalls(), uint64(0))
	})

	t.Run("ShouldReport config set to true logs", func(t *testing.T) {
		t.Parallel()

		config := buildNilConsumerConfig(true)
		mockLogger := loggermock.NewLoggerMock()
		progressReporter := progressreporters.NewLogSleepReporter(mockLogger, config.Progress)
		consumer := buildNilConsumer(config, mockLogger)

		go func() {
			progressReporter.ConsumerReport(consumer)
		}()

		time.Sleep(500 * time.Millisecond)
		progressReporter.Stop()

		require.Equal(t, mockLogger.GetPrintfCalls(), uint64(1))
	})
}

func TestLogSleepReporter_ProducerReport(t *testing.T) {
	t.Parallel()

	t.Run("ShouldReport config set to false does not log anything", func(t *testing.T) {
		t.Parallel()

		config := buildNilProducerConfig(false)
		mockLogger := loggermock.NewLoggerMock()
		progressReporter := progressreporters.NewLogSleepReporter(mockLogger, config.Progress)
		producer := buildNilProducer(config, mockLogger)

		progressReporter.ProducerReport(producer)

		require.Equal(t, mockLogger.GetPrintfCalls(), uint64(0))
	})

	t.Run("ShouldReport config set to true logs", func(t *testing.T) {
		t.Parallel()

		config := buildNilProducerConfig(true)
		mockLogger := loggermock.NewLoggerMock()
		progressReporter := progressreporters.NewLogSleepReporter(mockLogger, config.Progress)
		producer := buildNilProducer(config, mockLogger)

		go func() {
			progressReporter.ProducerReport(producer)
		}()

		time.Sleep(500 * time.Millisecond)
		progressReporter.Stop()

		require.Equal(t, mockLogger.GetPrintfCalls(), uint64(1))
	})
}

func TestLogSleepReporter_ShouldReportProgress(t *testing.T) {
	t.Parallel()

	t.Run("false config returns false", func(t *testing.T) {
		t.Parallel()

		config := buildNilProducerConfig(false)
		mockLogger := loggermock.NewLoggerMock()
		progressReporter := progressreporters.NewLogSleepReporter(mockLogger, config.Progress)

		shouldReport := progressReporter.ShouldReportProgress()

		require.False(t, shouldReport)
	})

	t.Run("true config returns true", func(t *testing.T) {
		t.Parallel()

		config := buildNilProducerConfig(true)
		mockLogger := loggermock.NewLoggerMock()
		progressReporter := progressreporters.NewLogSleepReporter(mockLogger, config.Progress)

		shouldReport := progressReporter.ShouldReportProgress()

		require.True(t, shouldReport)
	})
}

func TestLogSleepReporter_Stop(t *testing.T) {
	t.Parallel()

	config := buildNilProducerConfig(true)
	mockLogger := loggermock.NewLoggerMock()
	progressReporter := progressreporters.NewLogSleepReporter(mockLogger, config.Progress)

	shouldReport := progressReporter.ShouldReportProgress()
	require.True(t, shouldReport)

	progressReporter.Stop()
	shouldReport = progressReporter.ShouldReportProgress()

	require.False(t, shouldReport)
}

func buildNilConsumer(config *configs.SolverConfig, logger loggers.Logger) *nilconsumer.NilConsumer {
	dataChannel := make(<-chan *models.TemperatureReading)
	doneChannel := make(<-chan bool)
	return nilconsumer.NewNilConsumer(dataChannel, doneChannel, logger, config)
}

func buildNilConsumerConfig(shouldReportProgress bool) *configs.SolverConfig {
	return &configs.SolverConfig{
		Progress: &configs.ProgressSolverConfig{
			ShouldReport:    shouldReportProgress,
			SleepDurationMs: 1000,
		},
	}
}

func buildNilProducer(config *configs.SolverConfig, logger loggers.Logger) *nilproducer.NilProducer {
	dataChannel := make(chan<- *models.TemperatureReading)
	doneChannel := make(chan<- bool)
	csvReader := csvreadermock.NewCSVReaderMock()
	return nilproducer.NewNilProducer(csvReader, dataChannel, doneChannel, logger, config)
}

func buildNilProducerConfig(shouldReportProgress bool) *configs.SolverConfig {
	return &configs.SolverConfig{
		Progress: &configs.ProgressSolverConfig{
			ShouldReport:    shouldReportProgress,
			SleepDurationMs: 1000,
		},
	}
}
