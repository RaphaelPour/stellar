package hack

import "unsafe"

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Bool[T Number](in T) bool {
	return in != 0
}

func Wormhole(in bool) int {
	return (*(*int)(unsafe.Pointer(&in))) & int(1)
}
