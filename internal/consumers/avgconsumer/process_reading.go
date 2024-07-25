package avgconsumer

import "github.com/thtg88/1brc/internal/models"

func (atc *AvgTempConsumer) ProcessReading(reading *models.TemperatureReading) {
	if atc.Config.Debug {
		atc.Logger.Printf("consuming %v", reading)
	}

	defer func() { atc.RecordsConsumed++ }()

	cityStats, ok := atc.Stats[reading.City]
	if !ok {
		atc.Stats[reading.City] = atc.processEmptyReading(reading)

		return
	}

	atc.Stats[reading.City] = atc.processCumulativeReading(cityStats, reading)
}

func (atc *AvgTempConsumer) processEmptyReading(reading *models.TemperatureReading) models.CityStats {
	if atc.Config.CalculateAverageForEachReading {
		return models.CityStats{
			City:              reading.City,
			AverageTemp:       reading.Temperature,
			MeasurementsSum:   reading.Temperature,
			MeasurementsCount: 1,
		}
	}

	return models.CityStats{
		City:              reading.City,
		MeasurementsSum:   reading.Temperature,
		MeasurementsCount: 1,
	}
}

func (atc *AvgTempConsumer) processCumulativeReading(cityStats models.CityStats, reading *models.TemperatureReading) models.CityStats {
	sum := cityStats.MeasurementsSum + reading.Temperature
	count := cityStats.MeasurementsCount + 1
	if atc.Config.CalculateAverageForEachReading {
		return models.CityStats{
			City:              reading.City,
			AverageTemp:       sum / count,
			MeasurementsSum:   sum,
			MeasurementsCount: count,
		}
	}

	return models.CityStats{
		City:              reading.City,
		MeasurementsSum:   sum,
		MeasurementsCount: count,
	}
}
