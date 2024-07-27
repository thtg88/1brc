package rawfilereadsolver

import (
	"os"

	"github.com/thtg88/1brc/internal/consumers/bufferedsequentialconsumer"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/producers/rawfilereadproducer"
)

func (rfrs *RawFileReadSolver) ProcessTemperatures(file *os.File) []models.CityStats {
	dataChannel := make(chan []*models.TemperatureReading, 1)
	doneChannel := make(chan bool)

	consumer := bufferedsequentialconsumer.NewBufferedSequentialConsumer(dataChannel, doneChannel, rfrs.Logger, rfrs.Config)
	producer := rawfilereadproducer.NewRawFileReadProducer(file, dataChannel, doneChannel, rfrs.Logger, rfrs.Config)

	go rfrs.ProgressReporter.ProducerReport(producer)
	go rfrs.ProgressReporter.ConsumerReport(consumer)
	go producer.Start()

	consumer.Start()

	rfrs.StopProgressReport()

	return consumer.GetSortedStats()
}
