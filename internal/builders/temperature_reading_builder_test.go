package builders_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thtg88/1brc/internal/builders"
)

func TestTemperatureReadingBuilder_Build(t *testing.T) {
	t.Parallel()

	builder := builders.NewTemperatureReadingBuilder()
	reading := builder.Build()

	assert.Equal(t, "", reading.City)
	require.Zero(t, reading.Temperature)
}

func TestTemperatureReadingBuilder_WithCity(t *testing.T) {
	t.Parallel()

	const anotherCity = "another city"
	builder := builders.NewTemperatureReadingBuilder().WithCity(anotherCity)
	reading := builder.Build()

	assert.Equal(t, anotherCity, reading.City)
	require.Zero(t, reading.Temperature)
}

func TestTemperatureReadingBuilder_WithTemperature(t *testing.T) {
	t.Parallel()

	const temperature = int64(456)
	builder := builders.NewTemperatureReadingBuilder().WithTemperature(temperature)
	reading := builder.Build()

	assert.Equal(t, "", reading.City)
	assert.Equal(t, temperature, reading.Temperature)
}

func TestTemperatureReadingBuilder_WithTestValues(t *testing.T) {
	t.Parallel()

	builder := builders.NewTemperatureReadingBuilder().WithTestValues()
	reading := builder.Build()

	assert.Equal(t, builders.TemperatureReadingBuilder_TestCity, reading.City)
	assert.Equal(t, builders.TemperatureReadingBuilder_TestTemperature, reading.Temperature)
}
