package runes_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/runes"
)

func TestRuneAt(t *testing.T) {
	type test struct {
		descrtiption string
		str          string
		index        uint

		wantRune  rune
		wantFound bool
	}

	tests := []test{
		{
			descrtiption: "empty string returns not found",
			str:          "",
			index:        0,
			wantRune:     0,
			wantFound:    false,
		},
		{
			descrtiption: "string with an index too high returns not found",
			str:          "0123",
			index:        4,
			wantRune:     0,
			wantFound:    false,
		},
		{
			descrtiption: "rune at index found",
			str:          "0123",
			index:        0,
			wantRune:     '0',
			wantFound:    true,
		},
	}

	for _, testCase := range tests {
		tc := testCase
		t.Run(tc.descrtiption, func(t *testing.T) {
			t.Parallel()

			actualRune, actualFound := runes.RuneAt(tc.str, tc.index)

			require.Equal(t, tc.wantFound, actualFound)
			require.Equal(t, tc.wantRune, actualRune)
		})
	}
}
