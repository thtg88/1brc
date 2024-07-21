package progressreporters

import (
	"github.com/thtg88/1brc/internal/consumers"
	"github.com/thtg88/1brc/internal/producers"
)

type ProgressReporter interface {
	ConsumerReport(consumer consumers.Consumer)
	ProducerReport(producer producers.Producer)
	Stop()
}
