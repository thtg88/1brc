package rawfilereadsolver

import (
	"os"

	"github.com/thtg88/1brc/internal/consumers/bufferedsequentialconsumer"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/producers/rawfilereadproducer"
)

func (rfrs *RawFileReadSolver) ProcessTemperatures(file *os.File) ([]models.CityStats, error) {
	dataChannel := make(chan []*models.TemperatureReading, 1)

	consumer := bufferedsequentialconsumer.NewBufferedSequentialConsumer(dataChannel, rfrs.Logger, rfrs.Config)
	producer := rawfilereadproducer.NewRawFileReadProducer(file, dataChannel, rfrs.Logger, rfrs.Config)

	go rfrs.ProgressReporter.ProducerReport(producer)
	go rfrs.ProgressReporter.ConsumerReport(consumer)
	go producer.Start()

	consumer.Start()

	rfrs.StopProgressReport()

	return consumer.GetSortedStats(), nil
}
