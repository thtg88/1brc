package models

type CityStats struct {
	City              string
	MinTemp           int64
	MaxTemp           int64
	AverageTemp       int64
	MeasurementsSum   int64
	MeasurementsCount int64
}

func NewReadingCityStats(reading *TemperatureReading) CityStats {
	return CityStats{
		City:              reading.City,
		MinTemp:           reading.Temperature,
		MaxTemp:           reading.Temperature,
		AverageTemp:       reading.Temperature,
		MeasurementsSum:   reading.Temperature,
		MeasurementsCount: 1,
	}
}

func NewCumulativeReadingCityStats(stats CityStats, reading *TemperatureReading) CityStats {
	sum := stats.MeasurementsSum + reading.Temperature
	count := stats.MeasurementsCount + 1

	return CityStats{
		City:              reading.City,
		MinTemp:           min(reading.Temperature, stats.MinTemp),
		MaxTemp:           max(reading.Temperature, stats.MaxTemp),
		AverageTemp:       sum / count,
		MeasurementsSum:   sum,
		MeasurementsCount: count,
	}
}

func NewCumulativeReadingCityStatsWithoutAverage(stats CityStats, reading *TemperatureReading) CityStats {
	sum := stats.MeasurementsSum + reading.Temperature
	count := stats.MeasurementsCount + 1

	return CityStats{
		City:    reading.City,
		MinTemp: min(reading.Temperature, stats.MinTemp),
		MaxTemp: max(reading.Temperature, stats.MaxTemp),
		// AverageTemp:       sum / count,
		MeasurementsSum:   sum,
		MeasurementsCount: count,
	}
}
