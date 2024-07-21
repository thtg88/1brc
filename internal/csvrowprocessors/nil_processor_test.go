package csvrowprocessors_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/csvrowprocessors"
)

func TestNilProcessor_Process(t *testing.T) {
	t.Parallel()

	processor := csvrowprocessors.NewNilProcessor()

	actualReading, err := processor.Process([]string{})

	require.NoError(t, err)
	require.Nil(t, actualReading)
}
