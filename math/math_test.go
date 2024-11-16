package math

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPow(t *testing.T) {
	require.Equal(t, int(4), Pow(int(2), int(2)))
	require.Equal(t, uint(4), Pow(uint(2), uint(2)))
	require.Equal(t, float32(4), Pow(float32(2), float32(2)))
	require.Equal(t, float64(4), Pow(float64(2), float64(2)))
}

func TestToF(t *testing.T) {
	result := float64(4)
	require.Equal(t, result, ToF(int(4)))
	require.Equal(t, result, ToF(uint(4)))
	require.Equal(t, result, ToF(float32(4)))
	require.Equal(t, result, ToF(float64(4)))
}

func TestAbs(t *testing.T) {
	require.Equal(t, uint(2), Abs(uint(2)))
	require.Equal(t, int(2), Abs(int(2)))
	require.Equal(t, int(2), Abs(int(-2)))
	require.Equal(t, float32(2), Abs(float32(2)))
	require.Equal(t, float32(2), Abs(float32(-2)))
	require.Equal(t, float64(2), Abs(float64(2)))
	require.Equal(t, float64(2), Abs(float64(-2)))
}

func TestMin(t *testing.T) {
	require.Equal(t, uint(1), Min([]uint{1, 2, 3}...))
	require.Equal(t, uint(1), Min([]uint{3, 2, 1}...))

	require.Equal(t, int(1), Min([]int{1, 2, 3}...))
	require.Equal(t, float32(1), Min([]float32{1, 2, 3}...))
	require.Equal(t, float64(1), Min([]float64{1, 2, 3}...))
}

func TestMax(t *testing.T) {
	require.Equal(t, uint(3), Max([]uint{1, 2, 3}...))
	require.Equal(t, uint(3), Max([]uint{3, 2, 1}...))

	require.Equal(t, int(3), Max([]int{1, 2, 3}...))
	require.Equal(t, float32(3), Max([]float32{1, 2, 3}...))
	require.Equal(t, float64(3), Max([]float64{1, 2, 3}...))
}

func TestSign(t *testing.T) {
	require.Equal(t, int(1), Sign(int(10)))
	require.Equal(t, int(-1), Sign(int(-10)))
	require.Equal(t, float32(1), Sign(float32(10)))
	require.Equal(t, float32(-1), Sign(float32(-10)))
	require.Equal(t, float64(1), Sign(float64(10)))
	require.Equal(t, float64(-1), Sign(float64(-10)))
}

func TestMinMax(t *testing.T) {
	{
		min, max := MinMax([]uint{1, 2, 3}...)
		require.Equal(t, min, uint(1))
		require.Equal(t, max, uint(3))
	}
	{
		min, max := MinMax([]int{1, 2, 3}...)
		require.Equal(t, min, int(1))
		require.Equal(t, max, int(3))
	}
	{
		min, max := MinMax([]float32{1, 2, 3}...)
		require.Equal(t, min, float32(1))
		require.Equal(t, max, float32(3))
	}
	{
		min, max := MinMax([]float64{1, 2, 3}...)
		require.Equal(t, min, float64(1))
		require.Equal(t, max, float64(3))
	}
}

func TestMaxN(t *testing.T) {
	for _, data := range []struct {
		inputArray  []int
		inputLength int
		expected    []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{5, 4, 3}},
		{[]int{3, 4, 5}, 3, []int{5, 4, 3}},
		{[]int{3, 4}, 3, nil},
		{[]int{3}, 3, nil},
		{[]int{}, 3, nil},
	} {
		require.Equal(t, data.expected, MaxN(data.inputArray, data.inputLength))
	}

	require.Equal(t, []float64{3.3, 3.2}, MaxN([]float64{3.0, 3.3, 3.15, 3.1, 3.2}, 2))
}

func TestMinN(t *testing.T) {
	for _, data := range []struct {
		inputArray  []int
		inputLength int
		expected    []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 3}},
		{[]int{3, 4, 5}, 3, []int{3, 4, 5}},
		{[]int{3, 4}, 3, nil},
		{[]int{3}, 3, nil},
		{[]int{}, 3, nil},
	} {
		require.Equal(t, data.expected, MinN(data.inputArray, data.inputLength))
	}

	require.Equal(t, []float64{3.0, 3.1}, MinN([]float64{3.0, 3.3, 3.15, 3.1, 3.2}, 2))
}

func TestWithin(t *testing.T) {
	require.True(t, Within(uint(2), uint(1), uint(3)))
	require.True(t, Within(uint(1), uint(1), uint(3)))
	require.True(t, Within(uint(3), uint(1), uint(3)))
	require.False(t, Within(uint(0), uint(1), uint(3)))

	require.True(t, Within(int(2), int(1), int(3)))
	require.True(t, Within(float32(2), float32(1), float32(3)))
	require.True(t, Within(float64(2), float64(1), float64(3)))
}

func TestGCD(t *testing.T) {
	require.Equal(t, 2, GCD(6, 8))
	require.Equal(t, 5, GCD(15, 20))
}

func TestLCM(t *testing.T) {
	require.Equal(t, 10, LCM(2, 5))
}
