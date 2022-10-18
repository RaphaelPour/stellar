package strings

import (
	"fmt"
	"strconv"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(fmt.Sprintf("error converting %s to int: %s", in, err))
	}
	return out
}
