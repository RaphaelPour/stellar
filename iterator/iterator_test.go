package iterator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIteratorEachDigit(t *testing.T) {
	expected := []int{1, 3, 3, 7}
	actual := make([]int, 0)
	for digit := range EachDigit(1337, 10) {
		actual = append(actual, digit)
	}

	require.Equal(t, expected, actual)
}
