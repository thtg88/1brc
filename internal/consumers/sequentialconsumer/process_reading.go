package sequentialconsumer

import "github.com/thtg88/1brc/internal/models"

func (sc *SequentialConsumer) ProcessReading(reading *models.TemperatureReading) {
	if sc.Config.Debug {
		sc.Logger.Printf("consuming %v", reading)
	}

	defer func() { sc.RecordsConsumed++ }()

	cityStats, ok := sc.Stats[reading.City]
	if !ok {
		sc.Stats[reading.City] = models.NewReadingCityStats(reading)

		return
	}

	sc.Stats[reading.City] = models.NewCumulativeReadingCityStats(cityStats, reading)
}
