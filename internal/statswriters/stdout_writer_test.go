package statswriters_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/mocks/iowritermock"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/statswriters"
)

func TestStdoutWriter_Write(t *testing.T) {
	t.Parallel()

	type test struct {
		descrtiption string
		lines        []models.CityStats
	}

	tests := []test{
		{
			descrtiption: "writing an empty array returns 1 line (header)",
			lines:        []models.CityStats{},
		},
		{
			descrtiption: "writing n stats returns n+1 line",
			lines: []models.CityStats{{
				City:              "test city",
				MinTemp:           1,
				MaxTemp:           3,
				AverageTemp:       2,
				MeasurementsSum:   4,
				MeasurementsCount: 2,
			}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.descrtiption, func(t *testing.T) {
			t.Parallel()

			file := iowritermock.NewIOWriterMock()
			csvStatsWriter := statswriters.NewStdoutWriter(file)
			csvStatsWriter.Write(tc.lines)

			contents := string(file.GetBytes())

			expectedContents := "{"
			for _, row := range tc.lines {
				expectedContents = expectedContents + fmt.Sprintf(
					statswriters.CityStatsStdoutPattern,
					row.City,
					strconv.FormatFloat(float64(row.MinTemp)/10, 'f', 1, 64),
					strconv.FormatFloat(float64(row.MaxTemp)/10, 'f', 1, 64),
					strconv.FormatFloat(float64(row.AverageTemp)/10, 'f', 1, 64),
				)
			}
			expectedContents = expectedContents + "}"

			require.Equal(t, expectedContents, contents)
		})
	}
}
