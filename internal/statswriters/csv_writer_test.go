package statswriters_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/mocks/iowritermock"
	"github.com/thtg88/1brc/internal/models"
	"github.com/thtg88/1brc/internal/statswriters"
)

const (
	expectedCSVStringRowPattern = "%s,%s,%s,%s\n"
	expectedCSVHeaderRow        = "city,min,max,avg\n"
)

func TestCSVWriter_Write(t *testing.T) {
	t.Parallel()

	type test struct {
		descrtiption string
		lines        []models.CityStats

		wantLinesCount int
	}

	tests := []test{
		{
			descrtiption:   "writing an empty array returns 1 line (header)",
			lines:          []models.CityStats{},
			wantLinesCount: 1,
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
			wantLinesCount: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.descrtiption, func(t *testing.T) {
			t.Parallel()

			file := iowritermock.NewIOWriterMock()
			csvStatsWriter := statswriters.NewCSVWriter(file)
			csvStatsWriter.Write(tc.lines)

			contents := string(file.GetBytes())
			actualLinesCount := strings.Count(contents, "\n")

			require.Equal(t, tc.wantLinesCount, actualLinesCount)

			require.Contains(t, contents, expectedCSVHeaderRow)

			for _, row := range tc.lines {
				expectedStringRow := fmt.Sprintf(
					expectedCSVStringRowPattern,
					row.City,
					strconv.FormatFloat(float64(row.MinTemp)/10, 'f', 1, 64),
					strconv.FormatFloat(float64(row.MaxTemp)/10, 'f', 1, 64),
					strconv.FormatFloat(float64(row.AverageTemp)/10, 'f', 1, 64),
				)

				require.Contains(t, contents, expectedStringRow)
			}
		})
	}
}
