package rawfilereadproducer_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/builders"
	"github.com/thtg88/1brc/internal/mocks/loggermock"
	"github.com/thtg88/1brc/internal/producers/rawfilereadproducer"
)

func TestRawFileReadProducer_ReadRows(t *testing.T) {
	t.Parallel()

	t.Run("successful read", func(t *testing.T) {
		t.Parallel()

		type test struct {
			description       string
			bytesCount        uint64
			initialCarryover  string
			initialPosition   uint64
			expectedRows      [][]string
			expectedCarryover string
		}

		tests := []test{
			// no carryover
			{
				description:       "0 byte read, no initial carryover",
				bytesCount:        0,
				initialCarryover:  "",
				initialPosition:   0,
				expectedRows:      [][]string{},
				expectedCarryover: "",
			},
			{
				description:       "1 byte read, no initial carryover",
				bytesCount:        1,
				initialCarryover:  "",
				initialPosition:   0,
				expectedRows:      [][]string{},
				expectedCarryover: string(builders.TemperatureReadingBuilder_TestCity[0]),
			},
			{
				description:       "10 bytes read, no initial carryover",
				bytesCount:        10,
				initialCarryover:  "",
				initialPosition:   0,
				expectedRows:      [][]string{},
				expectedCarryover: fmt.Sprintf("%s%s", builders.TemperatureReadingBuilder_TestCity, rawfilereadproducer.CommaSeparator),
			},
			{
				// 1 row with 1 character of carryover
				description:      "16 bytes read, no initial carryover",
				bytesCount:       16,
				initialCarryover: "",
				initialPosition:  0,
				expectedRows: [][]string{
					{builders.TemperatureReadingBuilder_TestCity, builders.TemperatureReadingCSVRowBuilder_TestTemperature},
				},
				expectedCarryover: string(builders.TemperatureReadingBuilder_TestCity[0]),
			},
			{
				description:      "1000 bytes read, no initial carryover",
				bytesCount:       1000,
				initialCarryover: "",
				initialPosition:  0,
				expectedRows: [][]string{
					{builders.TemperatureReadingBuilder_TestCity, builders.TemperatureReadingCSVRowBuilder_TestTemperature},
					{builders.TemperatureReadingBuilder_TestCity, builders.TemperatureReadingCSVRowBuilder_TestTemperature},
				},
				expectedCarryover: "",
			},
			// carryover
			{
				description:       "0 byte read, with carryover",
				bytesCount:        0,
				initialCarryover:  string(builders.TemperatureReadingBuilder_TestCity[0]),
				initialPosition:   1,
				expectedRows:      [][]string{},
				expectedCarryover: string(builders.TemperatureReadingBuilder_TestCity[0]),
			},
			{
				description:      "1 byte read, with carryover",
				bytesCount:       1,
				initialCarryover: string(builders.TemperatureReadingBuilder_TestCity[0]),
				initialPosition:  1,
				expectedRows:     [][]string{},
				// `te`
				expectedCarryover: builders.TemperatureReadingBuilder_TestCity[:2],
			},
			{
				description:      "10 bytes read, with carryover",
				bytesCount:       10,
				initialCarryover: string(builders.TemperatureReadingBuilder_TestCity[0]),
				initialPosition:  1,
				expectedRows:     [][]string{},
				// c1234567890
				// test city,12.3
				expectedCarryover: fmt.Sprintf(
					"%s%s%s",
					builders.TemperatureReadingBuilder_TestCity,
					rawfilereadproducer.CommaSeparator,
					string(builders.TemperatureReadingCSVRowBuilder_TestTemperature[0]),
				),
			},
			{
				// 1 row with 1 character of carryover
				description:      "16 bytes read, with carryover",
				bytesCount:       16,
				initialCarryover: string(builders.TemperatureReadingBuilder_TestCity[0]),
				initialPosition:  1,
				// c1234567890123456
				// test city,12.3nte
				expectedRows: [][]string{
					{builders.TemperatureReadingBuilder_TestCity, builders.TemperatureReadingCSVRowBuilder_TestTemperature},
				},
				// `te`
				expectedCarryover: builders.TemperatureReadingBuilder_TestCity[:2],
			},
			{
				description:      "1000 bytes read, with carryover",
				bytesCount:       1000,
				initialCarryover: string(builders.TemperatureReadingBuilder_TestCity[0]),
				initialPosition:  1,
				expectedRows: [][]string{
					{builders.TemperatureReadingBuilder_TestCity, builders.TemperatureReadingCSVRowBuilder_TestTemperature},
					{builders.TemperatureReadingBuilder_TestCity, builders.TemperatureReadingCSVRowBuilder_TestTemperature},
				},
				expectedCarryover: "",
			},
		}

		for _, testCase := range tests {
			tc := testCase
			t.Run(tc.description, func(t *testing.T) {
				t.Parallel()

				mockLogger := loggermock.NewLoggerMock()
				producer := buildRawFileReadProducer(mockLogger, tc.bytesCount, tc.initialPosition, false, false, false)

				actualRows, actualCarryover, err := producer.ReadRows(tc.initialCarryover)

				assert.NoError(t, err)
				require.Equal(t, tc.expectedRows, actualRows)
				require.Equal(t, tc.expectedCarryover, actualCarryover)
			})
		}
	})

	t.Run("error when reading from IO reader bubbles up", func(t *testing.T) {
		t.Parallel()

		mockLogger := loggermock.NewLoggerMock()
		producer := buildRawFileReadProducer(mockLogger, 30, 0, false, false, true)

		expectedErr := fmt.Errorf("could not read from CSV file: %v", errors.New("forced error on IO reader"))

		actualRows, actualCarryover, err := producer.ReadRows("")

		require.Error(t, err)
		require.Equal(t, expectedErr, err)
		require.Equal(t, [][]string{}, actualRows)
		require.Equal(t, "", actualCarryover)
	})

	t.Run("EOF returns empty", func(t *testing.T) {
		t.Parallel()

		mockLogger := loggermock.NewLoggerMock()
		producer := buildRawFileReadProducer(mockLogger, 30, 0, false, true, false)

		actualRows, actualCarryover, err := producer.ReadRows("")

		require.NoError(t, err)
		require.Equal(t, [][]string{}, actualRows)
		require.Equal(t, "", actualCarryover)
	})

	t.Run("debug logs", func(t *testing.T) {
		t.Parallel()

		mockLogger := loggermock.NewLoggerMock()

		producer := buildRawFileReadProducer(mockLogger, 1000, 0, true, false, false)

		producer.ReadRows("")

		require.Equal(t, uint64(3), mockLogger.GetPrintfCalls())
	})
}
