package iter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPairs2(t *testing.T) {
	require.Equal(t, [][]int{
		[]int{1, 2},
		[]int{3, 4},
	},
		Pairs([]int{1, 2, 3, 4}, 2),
	)
}

func TestPairs3(t *testing.T) {
	require.Equal(t, [][]int{
		[]int{1, 2, 3},
		[]int{4},
	},
		Pairs([]int{1, 2, 3, 4}, 3),
	)
}
