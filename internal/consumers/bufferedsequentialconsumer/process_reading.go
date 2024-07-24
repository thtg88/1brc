package bufferedsequentialconsumer

import "github.com/thtg88/1brc/internal/models"

func (bsc *BufferedSequentialConsumer) ProcessReading(reading *models.TemperatureReading) {
	if bsc.Config.Debug {
		bsc.Logger.Printf("consuming %v", reading)
	}

	defer func() { bsc.RecordsConsumed++ }()

	cityStats, ok := bsc.Stats[reading.City]
	if !ok {
		bsc.Stats[reading.City] = models.NewReadingCityStats(reading)

		return
	}

	bsc.Stats[reading.City] = models.NewCumulativeReadingCityStats(cityStats, reading)
}
