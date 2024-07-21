package csvrowprocessors_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/csvrowprocessors"
)

func TestIntTempParseProcessor_Process(t *testing.T) {
	t.Parallel()

	t.Run("successful row process", func(t *testing.T) {

		processor := csvrowprocessors.NewIntTempParseProcessor()
		row := []string{builders.TemperatureReadingBuilder_TestCity, "12.3"}

		expectedReading := builders.NewTemperatureReadingBuilder().WithTestValues().Build()

		actualReading, err := processor.Process(row)

		require.NoError(t, err)
		require.Equal(t, expectedReading, actualReading)
	})

	t.Run("empty row returns an error", func(t *testing.T) {

		processor := csvrowprocessors.NewIntTempParseProcessor()
		row := []string{}

		reading, err := processor.Process(row)

		require.Error(t, err)
		require.Equal(t, errors.New("row length not 2: 0"), err)
		require.Nil(t, reading)
	})

	t.Run("row with length 1 returns an error", func(t *testing.T) {

		processor := csvrowprocessors.NewIntTempParseProcessor()
		row := []string{"test"}

		reading, err := processor.Process(row)

		require.Error(t, err)
		require.Equal(t, errors.New("row length not 2: 1"), err)
		require.Nil(t, reading)
	})

	t.Run("row with length 3 returns an error", func(t *testing.T) {

		processor := csvrowprocessors.NewIntTempParseProcessor()
		row := []string{"test", "problematic", "row"}

		reading, err := processor.Process(row)

		require.Error(t, err)
		require.Equal(t, errors.New("row length not 2: 3"), err)
		require.Nil(t, reading)
	})

	t.Run("invalid temperature returns an error", func(t *testing.T) {

		processor := csvrowprocessors.NewIntTempParseProcessor()
		row := []string{builders.TemperatureReadingBuilder_TestCity, "not a temperature"}

		reading, err := processor.Process(row)

		require.Error(t, err)
		require.Equal(t, errors.New("could not parse temperature: strconv.ParseInt: parsing \"not a temperatue\": invalid syntax"), err)
		require.Nil(t, reading)
	})
}
