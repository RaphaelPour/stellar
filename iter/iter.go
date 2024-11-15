package iter

import (
	"iter"
	"slices"

	"github.com/RaphaelPour/stellar/math"
)

func Pairs[S ~[]T, T any](in S, num int) []S {
	return slices.Collect(PairsSeq(in, num))
}

func PairsSeq[S ~[]T, T any](in S, num int) iter.Seq[S] {
	return func(yield func(S) bool) {
		for i := 0; i < len(in); i += num {
			if !yield(in[i:math.Min(i+num, len(in))]) {
				return
			}
		}
	}
}

func PairsSeq2[S ~[]T, T any](in S, num int) iter.Seq2[int, S] {
	return func(yield func(int, S) bool) {
		index := 0
		for i := 0; i < len(in); i += num {
			if !yield(index, in[i:math.Min(i+num, len(in))]) {
				return
			}
			index += 1
		}
	}
}
