package math

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPoint(t *testing.T) {
	p := Point[int]{1, 2}
	require.Equal(t, 1, p.X)
	require.Equal(t, 2, p.Y)

	newP := p.Add(Point[int]{2, 1})
	require.Equal(t, 3, newP.X)
	require.Equal(t, 3, newP.Y)

	// old point shouldn't have changed
	require.Equal(t, 1, p.X)
	require.Equal(t, 2, p.Y)
}
