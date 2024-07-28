package multirawfilereadsolver

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/consumers/bufferedsequentialconsumer"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/producers/rawfilereadproducer"
)

func (mrfrs *MultiRawFileReadSolver) ProcessTemperatures(originFile *os.File) ([]models.CityStats, error) {
	producersCount := int64(runtime.NumCPU() - 1)
	consumers := make([]*bufferedsequentialconsumer.BufferedSequentialConsumer, producersCount)

	positions, fileSize, err := mrfrs.calculateProducersPositions(originFile, producersCount)
	if err != nil {
		return []models.CityStats{}, fmt.Errorf("could not calculate producers positions: %v", err)
	}

	var wg sync.WaitGroup
	for producerNumber := int64(0); producerNumber < producersCount; producerNumber++ {
		file, err := mrfrs.openAndSeekFile(positions[producerNumber])
		if err != nil {
			return []models.CityStats{}, fmt.Errorf("could not open and seek file: %v", err)
		}
		defer file.Close()

		mrfrs.Logger.Printf("starting producer and consumer %d", producerNumber)

		dataChannel := make(chan []*models.TemperatureReading, 1)
		doneChannel := make(chan bool)

		consumers[producerNumber] = bufferedsequentialconsumer.NewBufferedSequentialConsumer(dataChannel, doneChannel, mrfrs.Logger, mrfrs.Config)

		producerConfig := configs.NewDefaultSolverConfig()
		producerConfig.FilePositioning.Enabled = true
		if producerNumber == producersCount-1 {
			// If last producer, read until the end
			producerConfig.FilePositioning.ReadUntilFilePosition = fileSize
		} else {
			producerConfig.FilePositioning.ReadUntilFilePosition = positions[producerNumber+1]
		}

		producer := rawfilereadproducer.NewRawFileReadProducer(file, dataChannel, doneChannel, mrfrs.Logger, producerConfig)

		go mrfrs.ProgressReporter.ProducerReport(producer)
		go mrfrs.ProgressReporter.ConsumerReport(consumers[producerNumber])

		wg.Add(2)
		go func(producer *rawfilereadproducer.RawFileReadProducer) {
			defer wg.Done()

			producer.Start()
		}(producer)
		go func(consumer *bufferedsequentialconsumer.BufferedSequentialConsumer) {
			defer wg.Done()

			consumer.Start()
		}(consumers[producerNumber])
	}

	wg.Wait()

	mrfrs.StopProgressReport()

	return mrfrs.combineConsumerStats(consumers)
}

func (mrfrs *MultiRawFileReadSolver) calculateProducersPositions(originFile *os.File, producersCount int64) ([]int64, int64, error) {
	fileInfo, err := originFile.Stat()
	if err != nil {
		return []int64{}, 0, fmt.Errorf("could not stat: %v", err)
	}

	fileSize := fileInfo.Size()
	positions := make([]int64, producersCount)
	for producerNumber := int64(0); producerNumber < producersCount; producerNumber++ {
		startingPosition := producerNumber * (fileSize / producersCount)

		positions[producerNumber], err = mrfrs.seekToNewLine(originFile, startingPosition)
		if err != nil {
			return []int64{}, fileSize, fmt.Errorf("could not seek to new line: %v", err)
		}
	}

	return positions, fileSize, nil
}

func (mrfrs *MultiRawFileReadSolver) openAndSeekFile(position int64) (*os.File, error) {
	file, err := os.Open(mrfrs.Config.SourceFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", mrfrs.Config.SourceFilePath, err)
	}

	_, err = file.Seek(position, io.SeekStart)
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("could not seek: %v", err)
	}

	return file, nil
}

func (mrfrs *MultiRawFileReadSolver) seekToNewLine(file *os.File, currentPosition int64) (int64, error) {
	if currentPosition == 0 {
		return currentPosition, nil
	}

	_, err := file.Seek(currentPosition, io.SeekStart)
	if err != nil {
		return currentPosition, fmt.Errorf("could not seek: %v", err)
	}

	bytes := make([]byte, 1)

	for string(bytes) != "\n" {
		_, err = file.Read(bytes)
		if err != nil {
			return currentPosition, fmt.Errorf("could not read: %v", err)
		}
		if err == io.EOF {
			return currentPosition, errors.New("EOF reached")
		}

		currentPosition++

		_, err := file.Seek(currentPosition, io.SeekStart)
		if err != nil {
			return currentPosition, fmt.Errorf("could not seek: %v", err)
		}
	}

	return currentPosition, nil
}

func (mrfrs *MultiRawFileReadSolver) combineConsumerStats(consumers []*bufferedsequentialconsumer.BufferedSequentialConsumer) ([]models.CityStats, error) {
	stats := []models.CityStats{}
	for _, consumer := range consumers {
		sortedStats := consumer.GetSortedStats()
		if len(stats) == 0 {
			stats = append(stats, sortedStats...)
			continue
		}

		// TODO: this may not work at lower numbers of records consumed,
		// as not all consumers may have consumed temperature data about all cities
		for idx, cityStats := range stats {
			consumerCityStats, err := consumer.GetCityStats(cityStats.City)
			if err != nil {
				return []models.CityStats{}, err
			}

			sum := cityStats.MeasurementsSum + consumerCityStats.MeasurementsSum
			count := cityStats.MeasurementsCount + consumerCityStats.MeasurementsCount
			stats[idx] = models.CityStats{
				City:              cityStats.City,
				MinTemp:           min(cityStats.MinTemp, consumerCityStats.MinTemp),
				MaxTemp:           max(cityStats.MaxTemp, consumerCityStats.MaxTemp),
				AverageTemp:       sum / count,
				MeasurementsSum:   sum,
				MeasurementsCount: count,
			}
		}
	}

	return stats, nil
}
