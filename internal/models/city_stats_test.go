package models_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/models"
)

func TestNewReadingCityStats(t *testing.T) {
	t.Parallel()

	reading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()
	wantCityStats := models.CityStats{
		City:              builders.TemperatureReadingBuilder_TestCity,
		MinTemp:           builders.TemperatureReadingBuilder_TestTemperature,
		MaxTemp:           builders.TemperatureReadingBuilder_TestTemperature,
		AverageTemp:       builders.TemperatureReadingBuilder_TestTemperature,
		MeasurementsSum:   builders.TemperatureReadingBuilder_TestTemperature,
		MeasurementsCount: 1,
	}

	actualCityStats := models.NewReadingCityStats(reading)

	require.Equal(t, wantCityStats, actualCityStats)
}

func TestNewCumulativeReadingCityStats(t *testing.T) {
	t.Parallel()

	existingCityStats := models.CityStats{
		City:              "test city",
		MinTemp:           3,
		MaxTemp:           3,
		AverageTemp:       3,
		MeasurementsSum:   3,
		MeasurementsCount: 1,
	}
	reading := &models.TemperatureReading{
		City:        "test city",
		Temperature: 1,
	}
	wantCityStats := models.CityStats{
		City:              "test city",
		MinTemp:           1,
		MaxTemp:           3,
		AverageTemp:       2,
		MeasurementsSum:   4,
		MeasurementsCount: 2,
	}

	actualCityStats := models.NewCumulativeReadingCityStats(existingCityStats, reading)

	require.Equal(t, wantCityStats, actualCityStats)
}

func TestNewCumulativeReadingCityStatsWithoutAverage(t *testing.T) {
	t.Parallel()

	existingCityStats := models.CityStats{
		City:              "test city",
		MinTemp:           3,
		MaxTemp:           3,
		AverageTemp:       3,
		MeasurementsSum:   3,
		MeasurementsCount: 1,
	}
	reading := &models.TemperatureReading{
		City:        "test city",
		Temperature: 1,
	}
	wantCityStats := models.CityStats{
		City:              "test city",
		MinTemp:           1,
		MaxTemp:           3,
		MeasurementsSum:   4,
		MeasurementsCount: 2,
	}

	actualCityStats := models.NewCumulativeReadingCityStatsWithoutAverage(existingCityStats, reading)

	require.Equal(t, wantCityStats, actualCityStats)
}
