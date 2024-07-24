package builders_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
)

func TestTemperatureReadingCSVRowBuilder_Build(t *testing.T) {
	t.Parallel()

	builder := builders.NewTemperatureReadingCSVRowBuilder()
	row := builder.Build()

	assert.Len(t, row, builders.CSVRowLength)
	require.Equal(t, "", row[0])
	require.Equal(t, "", row[1])
}

func TestTemperatureReadingCSVRowBuilder_WithTestValues(t *testing.T) {
	t.Parallel()

	builder := builders.NewTemperatureReadingCSVRowBuilder().WithTestValues()
	row := builder.Build()

	assert.Len(t, row, builders.CSVRowLength)
	require.Equal(t, builders.TemperatureReadingBuilder_TestCity, row[0])
	require.Equal(t, builders.TemperatureReadingCSVRowBuilder_TestTemperature, row[1])
}
