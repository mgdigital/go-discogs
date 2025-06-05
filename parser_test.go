package discogs

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseDiscTrackNumber(t *testing.T) {
	for _, test := range []struct {
		input       string
		discNumber  int
		trackNumber int
		err         error
	}{
		{
			input:       "1",
			discNumber:  1,
			trackNumber: 1,
		},
		{
			input:       "2",
			discNumber:  1,
			trackNumber: 2,
		},
		{
			input:       "11",
			discNumber:  1,
			trackNumber: 11,
		},
		{
			input:       "A",
			discNumber:  1,
			trackNumber: 1,
		},
		{
			input:       "B",
			discNumber:  2,
			trackNumber: 1,
		},
		{
			input:       "B3",
			discNumber:  2,
			trackNumber: 3,
		},
		{
			input:       "3-4",
			discNumber:  3,
			trackNumber: 4,
		},
	} {
		t.Run(test.input, func(t *testing.T) {
			discNumber, trackNumber, err := ParseDiscTrackNumber(test.input)

			if err != nil {
				require.ErrorIs(t, err, test.err)
			} else {
				assert.Equal(t, test.discNumber, discNumber)
				assert.Equal(t, test.trackNumber, trackNumber)
			}
		})
	}
}
