package math

func Product[T Number](vector []T) T {
	if len(vector) == 0 {
		return 0
	}

	product := T(1)
	for _, factor := range vector {
		product *= factor
	}
	return product
}

func Sum[T Number](vector []T) T {
	sum := T(0)
	for _, summand := range vector {
		sum += summand
	}
	return sum
}
