package generator

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

func TestReverse(t *testing.T) {
	require.Equal(t, []int{3, 2, 1}, Reverse([]int{1, 2, 3}))
}

func TestSkipOne(t *testing.T) {
	require.Equal(
		t,
		[][]int{
			[]int{2, 3},
			[]int{1, 3},
			[]int{1, 2},
		},
		SkipOne([]int{1, 2, 3}),
	)
}
