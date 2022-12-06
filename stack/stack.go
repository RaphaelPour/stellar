package stack

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) PushAhead(item T) {
	s.items = append([]T{item}, s.items...)
}

func (s *Stack[T]) PushN(items []T) {
	s.items = append(s.items, items...)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) PopN(l int) []T {
	if len(s.items) < l {
		panic("stack has not enough items")
	}
	items := append([]T{}, s.items[len(s.items)-l:]...)
	s.items = s.items[:len(s.items)-l]
	return items
}

func (s *Stack[T]) Clear() {
	s.items = nil
}
