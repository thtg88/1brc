package parsers_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/parsers"
)

func TestParseTemperatureFloat(t *testing.T) {
	t.Parallel()

	actualTemperature, actualErr := parsers.ParseTemperatureFloat("123.4")

	require.Equal(t, 123.4, actualTemperature)
	require.Equal(t, nil, actualErr)
}

func TestParseTemperatureInt(t *testing.T) {
	t.Parallel()

	actualTemperature, actualErr := parsers.ParseTemperatureInt("123.4")

	require.Equal(t, int64(1234), actualTemperature)
	require.Equal(t, nil, actualErr)
}
