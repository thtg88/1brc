package csvrowparsers_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/csvrowparsers"
)

func TestNilParser_Parse(t *testing.T) {
	t.Parallel()

	parser := csvrowparsers.NewNilParser()

	actualReading, err := parser.Parse([]string{})

	require.NoError(t, err)
	require.Nil(t, actualReading)
}
