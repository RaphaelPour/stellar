package generator

import (
	"slices"
	"testing"

	"github.com/RaphaelPour/stellar/math"
	"github.com/stretchr/testify/require"
)

func TestDirections(t *testing.T) {
	require.Equal(t,
		[]math.Point{
			math.Point{X: -1, Y: -1},
			math.Point{X: 0, Y: -1},
			math.Point{X: 1, Y: -1},
			math.Point{X: -1, Y: 0},
			math.Point{X: 1, Y: 0},
			math.Point{X: -1, Y: 1},
			math.Point{X: 0, Y: 1},
			math.Point{X: 1, Y: 1},
		},
		slices.Collect(DirectionsSeq()),
	)
}

func TestDirectionsReverse(t *testing.T) {
	require.Equal(t,
		Reverse([]math.Point{
			math.Point{X: -1, Y: -1},
			math.Point{X: 0, Y: -1},
			math.Point{X: 1, Y: -1},
			math.Point{X: -1, Y: 0},
			math.Point{X: 1, Y: 0},
			math.Point{X: -1, Y: 1},
			math.Point{X: 0, Y: 1},
			math.Point{X: 1, Y: 1},
		}),
		slices.Collect(DirectionsSeq(WithReverse())),
	)
}
