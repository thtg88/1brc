package progressreporters

import (
	"time"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers"
	"github.com/thtg88/1brc/internal/loggers"
	"github.com/thtg88/1brc/internal/producers"
)

type LogSleepReporter struct {
	logger loggers.Logger
	config *configs.ProgressSolverConfig
}

const (
	ConsumerLogFormat = "consumed %d/1,000,000,000"
	ProducerLogFormat = "produced %d/1,000,000,000"
)

func NewLogSleepReporter(logger loggers.Logger, config *configs.ProgressSolverConfig) *LogSleepReporter {
	return &LogSleepReporter{
		logger: logger,
		config: config,
	}
}

func (lsr *LogSleepReporter) ConsumerReport(consumer consumers.Consumer) {
	for {
		if !lsr.Enabled() {
			return
		}

		lsr.logger.Printf(ConsumerLogFormat, consumer.GetRecordsConsumed())

		time.Sleep(time.Duration(lsr.config.SleepDurationMs) * time.Millisecond)
	}
}

func (lsr *LogSleepReporter) ProducerReport(producer producers.Producer) {
	for {
		if !lsr.Enabled() {
			return
		}

		lsr.logger.Printf(ProducerLogFormat, producer.GetRecordsProduced())

		time.Sleep(time.Duration(lsr.config.SleepDurationMs) * time.Millisecond)
	}
}

func (lsr *LogSleepReporter) Enabled() bool {
	return lsr.config.Enabled
}

func (lsr *LogSleepReporter) Stop() {
	lsr.config.Enabled = false
}
