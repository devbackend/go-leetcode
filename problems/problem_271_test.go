package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblemCodec(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		strs []string
	}{
		"example 1": {
			strs: []string{"Hello", "World", "!"},
		},
		"example 2": {
			strs: []string{""},
		},
		"example 3": {
			strs: []string{"Hello from", "ASCII"},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m1, m2 := Codec{}, Codec{}

			encoded := m1.Encode(tc.strs)

			require.Equal(t, tc.strs, m2.Decode(encoded))
		})
	}
}
