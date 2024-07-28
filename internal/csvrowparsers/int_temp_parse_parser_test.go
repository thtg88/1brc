package csvrowparsers_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/csvrowparsers"
)

func TestIntTempParser_Parse(t *testing.T) {
	t.Parallel()

	t.Run("successful row process", func(t *testing.T) {

		parser := csvrowparsers.NewIntTempParser()
		row := []string{builders.TemperatureReadingBuilder_TestCity, "12.3"}

		expectedReading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()

		actualReading, err := parser.Parse(row)

		require.NoError(t, err)
		require.Equal(t, expectedReading, actualReading)
	})

	t.Run("empty row returns an error", func(t *testing.T) {

		parser := csvrowparsers.NewIntTempParser()
		row := []string{}

		reading, err := parser.Parse(row)

		require.Nil(t, reading)
		assert.Error(t, err)
		require.Equal(t, errors.New("row length not 2 (0): []"), err)
	})

	t.Run("row with length 1 returns an error", func(t *testing.T) {

		parser := csvrowparsers.NewIntTempParser()
		row := []string{"test"}

		reading, err := parser.Parse(row)

		require.Nil(t, reading)
		assert.Error(t, err)
		require.Equal(t, errors.New("row length not 2 (1): [test]"), err)
	})

	t.Run("row with length 3 returns an error", func(t *testing.T) {

		parser := csvrowparsers.NewIntTempParser()
		row := []string{"test", "problematic", "row"}

		reading, err := parser.Parse(row)

		require.Nil(t, reading)
		assert.Error(t, err)
		require.Equal(t, errors.New("row length not 2 (3): [test problematic row]"), err)
	})

	t.Run("invalid temperature returns an error", func(t *testing.T) {

		parser := csvrowparsers.NewIntTempParser()
		row := []string{builders.TemperatureReadingBuilder_TestCity, "not a temperature"}

		expectedErr := errors.New("could not parse temperature: strconv.ParseInt: parsing \"not a temperatue\": invalid syntax")

		reading, err := parser.Parse(row)

		require.Nil(t, reading)
		assert.Error(t, err)
		require.Equal(t, expectedErr, err)
	})
}

func BenchmarkIntTempParser_Parse(b *testing.B) {
	parser := csvrowparsers.NewIntTempParser()
	row := []string{builders.TemperatureReadingBuilder_TestCity, "12.3"}
	for i := 0; i < b.N; i++ {
		parser.Parse(row)
	}
}
