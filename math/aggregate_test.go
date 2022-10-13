package math

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct(t *testing.T) {
	require.Equal(t, uint(42), Product([]uint{21, 2}))
	require.Equal(t, int(42), Product([]int{21, 2}))
	require.Equal(t, float32(42), Product([]float32{21, 2}))
	require.Equal(t, float64(42), Product([]float64{21, 2}))
}

func TestSum(t *testing.T) {
	require.Equal(t, uint(42), Sum([]uint{21, 21}))
	require.Equal(t, int(42), Sum([]int{21, 21}))
	require.Equal(t, float32(42), Sum([]float32{21, 21}))
	require.Equal(t, float64(42), Sum([]float64{21, 21}))
}
