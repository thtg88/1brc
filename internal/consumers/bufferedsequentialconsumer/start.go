package bufferedsequentialconsumer

import (
	"time"

	"github.com/thtg88/1brc/internal/models"
)

func (bsc *BufferedSequentialConsumer) Start() {
	for {
		select {
		case readings, more := <-bsc.DataChannel:
			if !more {
				bsc.Logger.Println("consumer channel closed")

				bsc.calculateAverages()

				return
			}

			for _, reading := range readings {
				bsc.ProcessReading(reading)
			}
		default:
			if bsc.Config.Debug {
				bsc.Logger.Println("waiting for readings to be produced")
			}

			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (bsc *BufferedSequentialConsumer) calculateAverages() {
	for city, cityStats := range bsc.Stats {
		bsc.Stats[city] = models.CityStats{
			City:              city,
			MinTemp:           cityStats.MinTemp,
			MaxTemp:           cityStats.MaxTemp,
			AverageTemp:       cityStats.MeasurementsSum / cityStats.MeasurementsCount,
			MeasurementsSum:   cityStats.MeasurementsSum,
			MeasurementsCount: cityStats.MeasurementsCount,
		}
	}
}
