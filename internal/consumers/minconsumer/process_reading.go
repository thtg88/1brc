package minconsumer

import "github.com/thtg88/1brc/internal/models"

func (mtc *MinTempConsumer) ProcessReading(reading *models.TemperatureReading) {
	if mtc.Config.Debug {
		mtc.Logger.Printf("consuming %v", reading)
	}

	defer func() { mtc.RecordsConsumed++ }()

	value, ok := mtc.Stats[reading.City]
	if !ok {
		mtc.Stats[reading.City] = reading.Temperature

		return
	}

	mtc.Stats[reading.City] = min(reading.Temperature, value)
}
