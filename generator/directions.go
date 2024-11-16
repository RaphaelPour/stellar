package generator

import (
	"iter"

	"github.com/RaphaelPour/stellar/math"
)

type directionOptions struct {
	orthogonalOnly bool
	reverse        bool
}

type DirectionOption func(options *directionOptions)

func WithOrthogonalOnly() DirectionOption {
	return func(options *directionOptions) {
		options.orthogonalOnly = true
	}
}

func WithReverse() DirectionOption {
	return func(options *directionOptions) {
		options.reverse = true
	}
}

func gte[T math.Number](a, b T) bool {
	return a >= b
}
func lte[T math.Number](a, b T) bool {
	return a <= b
}

func DirectionsSeq(options ...DirectionOption) iter.Seq[math.Point] {
	dopts := new(directionOptions)
	for _, option := range options {
		option(dopts)
	}

	return func(yield func(math.Point) bool) {
		start, step, end, cond := -1, 1, 1, lte[int]
		if dopts.reverse {
			start, step, end, cond = 1, -1, -1, gte[int]
		}
		for y := start; cond(y, end); y += step {
			for x := start; cond(x, end); x += step {
				if x == 0 && y == 0 {
					continue
				}

				if dopts.orthogonalOnly && x != 0 && y != 0 {
					continue
				}

				if !yield(math.Point{X: x, Y: y}) {
					return
				}
			}
		}
	}
}
