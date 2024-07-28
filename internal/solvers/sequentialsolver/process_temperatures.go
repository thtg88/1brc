package sequentialsolver

import (
	"encoding/csv"
	"os"

	"github.com/thtg88/1brc/internal/consumers/sequentialconsumer"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/producers/singleproducer"
)

func (ss *SequentialSolver) ProcessTemperatures(file *os.File) ([]models.CityStats, error) {
	dataChannel := make(chan *models.TemperatureReading, ss.Config.BufferedChannelSize)
	csvReader := csv.NewReader(file)

	consumer := sequentialconsumer.NewSequentialConsumer(dataChannel, ss.Logger, ss.Config)
	producer := singleproducer.NewSingleProducer(csvReader, dataChannel, ss.Logger, ss.Config)

	go ss.ProgressReporter.ProducerReport(producer)
	go ss.ProgressReporter.ConsumerReport(consumer)
	go producer.Start()

	consumer.Start()

	ss.StopProgressReport()

	return consumer.GetSortedStats(), nil
}
