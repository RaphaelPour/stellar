package iterator

import (
	"golang.org/x/exp/constraints"
)

func EachDigit[T constraints.Integer](input, base T) <-chan T {
	iteratorChannel := make(chan T)
	go func(input, base T) {
		for input > 0 {
			iteratorChannel <- input % base
			input /= base
		}
		close(iteratorChannel)
	}(input, base)
	return iteratorChannel
}
