package math

import (
	coreMath "math"
)

type Number interface {
	uint | Signed
}

type Signed interface {
	int | float32 | float64
}

func ToF[T Number](in T) float64 {
	return float64(in)
}

func Pow[T Number](base T, exponent T) T {
	return T(coreMath.Pow(ToF(base), ToF(exponent)))
}

func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Sign[T Signed](n T) T {
	if n < 0 {
		return T(-1)
	}
	return T(1)
}

func Min[T Number](nums []T) T {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func Max[T Number](nums []T) T {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func MinMax[T Number](nums []T) (T, T) {
	min := nums[0]
	max := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}
	return min, max
}

func Within[T Number](n, lowerBound, upperBound T) bool {
	return n >= lowerBound && n <= upperBound
}
