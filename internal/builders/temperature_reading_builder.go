package builders

import "github.com/thtg88/1brc/internal/models"

type TemperatureReadingBuilder struct {
	temperatureReading *models.TemperatureReading
}

const (
	TemperatureReadingBuilder_TestCity        = "test city"
	TemperatureReadingBuilder_TestTemperature = int64(123)
)

func NewTemperatureReadingBuilder() *TemperatureReadingBuilder {
	return &TemperatureReadingBuilder{temperatureReading: &models.TemperatureReading{}}
}

func (trb *TemperatureReadingBuilder) Build() *models.TemperatureReading {
	return trb.temperatureReading
}

func (trb *TemperatureReadingBuilder) WithCity(city string) *TemperatureReadingBuilder {
	trb.temperatureReading.City = city
	return trb
}

func (trb *TemperatureReadingBuilder) WithTemperature(temperature int64) *TemperatureReadingBuilder {
	trb.temperatureReading.Temperature = temperature
	return trb
}

func (trb *TemperatureReadingBuilder) WithTestValues() *TemperatureReadingBuilder {
	trb.temperatureReading.City = TemperatureReadingBuilder_TestCity
	trb.temperatureReading.Temperature = TemperatureReadingBuilder_TestTemperature
	return trb
}
