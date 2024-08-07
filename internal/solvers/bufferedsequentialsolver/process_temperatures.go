package bufferedsequentialsolver

import (
	"encoding/csv"
	"os"

	"github.com/thtg88/1brc/internal/consumers/bufferedsequentialconsumer"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/producers/bufferedproducer"
)

func (bss *BufferedSequentialSolver) ProcessTemperatures(file *os.File) ([]models.CityStats, error) {
	dataChannel := make(chan []*models.TemperatureReading, 1)
	csvReader := csv.NewReader(file)

	consumer := bufferedsequentialconsumer.NewBufferedSequentialConsumer(dataChannel, bss.Logger, bss.Config)
	producer := bufferedproducer.NewBufferedProducer(csvReader, dataChannel, bss.Logger, bss.Config)

	go bss.ProgressReporter.ProducerReport(producer)
	go bss.ProgressReporter.ConsumerReport(consumer)
	go producer.Start()

	consumer.Start()

	bss.StopProgressReport()

	return consumer.GetSortedStats(), nil
}
