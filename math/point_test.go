package math

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPoint(t *testing.T) {
	p := Point{1, 2}
	require.Equal(t, Point{1, 2}, p)

	newP := p.Add(Point{2, 1})
	require.Equal(t, Point{3, 3}, newP)

	// old point shouldn't have changed
	require.Equal(t, Point{1, 2}, p)

	require.Equal(t, "1/2", p.String())
}

func TestDistance(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{10, 10}
	require.Equal(t, 20, p1.ManhattanDist(p2))
	require.Equal(t, 14, int(p1.EuclideanDist(p2)))
}

func TestMinMaxPoint(t *testing.T) {
	p1 := Point{1, 2}
	p2 := Point{0, 3}

	require.Equal(t, Point{1, 3}, p1.Max(p2))
	require.Equal(t, Point{0, 2}, p1.Min(p2))
}
