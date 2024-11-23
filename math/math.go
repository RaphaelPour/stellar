package math

import (
	coreMath "math"
	"sort"

	"github.com/RaphaelPour/stellar/hack"
)

type Number interface {
	uint | Signed
}

type Signed interface {
	int | float32 | float64
}

type Integer interface {
	~int | ~int32 | ~int64 | ~uint
}

func ToF[T Number](in T) float64 {
	return float64(in)
}

func Pow[T Number](base T, exponent T) T {
	return T(coreMath.Pow(ToF(base), ToF(exponent)))
}

func Abs[T Number](n T) T {
	return -n*T(hack.Wormhole(n < 0)) + n*T(hack.Wormhole(n >= 0))
}

func Sign[T Signed](n T) T {
	return T(-1)*T(hack.Wormhole(n < 0)) + T(1)*T(hack.Wormhole(n >= 0))
}

func Min[T Number](nums ...T) T {
	return MinTraditional[T](nums...)
}

func MinBranchless[T Number](nums ...T) T {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		min = min*T(hack.Wormhole(nums[i] >= min)) + nums[i]*T(hack.Wormhole(nums[i] < min))
	}
	return min
}

func MinTraditional[T Number](nums ...T) T {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func Max[T Number](nums ...T) T {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func MinMax[T Number](nums ...T) (T, T) {
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

func MaxN[T Number](nums []T, n int) []T {
	if len(nums) < n {
		return nil
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	return nums[0:n]
}

func MinN[T Number](nums []T, n int) []T {
	if len(nums) < n {
		return nil
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return nums[0:n]
}

func Within[T Number](n, lowerBound, upperBound T) bool {
	return n >= lowerBound && n <= upperBound
}

func GCD[T Integer](a, b T) T {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM[T Integer](a, b T) T {
	return a * b / GCD(a, b)
}
