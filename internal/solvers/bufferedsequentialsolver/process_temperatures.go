package bufferedsequentialsolver

import (
	"encoding/csv"
	"os"

	"github.com/thtg88/1brc/internal/consumers/bufferedsequentialconsumer"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/producers/bufferedproducer"
	"github.com/thtg88/1brc/internal/progressreporters"
)

func (bss *BufferedSequentialSolver) ProcessTemperatures(file *os.File) []models.CityStats {
	dataChannel := make(chan []*models.TemperatureReading, 1)
	doneChannel := make(chan bool)
	csvReader := csv.NewReader(file)

	consumer := bufferedsequentialconsumer.NewBufferedSequentialConsumer(dataChannel, doneChannel, bss.Logger, bss.Config)
	producer := bufferedproducer.NewBufferedProducer(csvReader, dataChannel, doneChannel, bss.Logger, bss.Config)

	progressReporter := progressreporters.NewLogSleepReporter(bss.Logger, bss.Config.Progress)

	go progressReporter.ProducerReport(producer)
	go progressReporter.ConsumerReport(consumer)
	go producer.Start()

	consumer.Start()

	return consumer.GetSortedStats()
}
