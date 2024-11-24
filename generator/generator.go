package generator

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

func Reverse[S ~[]T, T any](in S) S {
	return slices.Collect(ReverseSeq(in))
}

func ReverseSeq[S ~[]T, T any](in S) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(in) - 1; i >= 0; i-- {
			if !yield(in[i]) {
				return
			}
		}
	}
}

func ReverseSeq2[S ~[]T, T any](in S) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := len(in) - 1; i >= 0; i-- {
			if !yield(i-len(in)-1, in[i]) {
				return
			}
		}
	}
}

func Ring[S ~[]T, T any](in S) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; ; i = (i + 1) % len(in) {
			if !yield(in[i]) {
				return
			}
		}
	}
}

func Ring2[S ~[]T, T any](in S) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := 0; ; i = (i + 1) % len(in) {
			if !yield(i, in[i]) {
				return
			}
		}
	}
}

func RingReverse[S ~[]T, T any](in S) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; ; i = (i + 1) % len(in) {
			if !yield(in[len(in)-i-1]) {
				return
			}
		}
	}
}

func RingReverse2[S ~[]T, T any](in S) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := 0; ; i = (i + 1) % len(in) {
			if !yield(len(in)-i-1, in[len(in)-i-1]) {
				return
			}
		}
	}
}
