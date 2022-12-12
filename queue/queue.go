package queue

type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	return make([]T, 0)
}

func (q *Queue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

func (q *Queue[T]) Dequeue() T {
	if len(*q) == 0 {
		panic("queue is empty")
	}

	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q Queue[T]) Count() int {
	return len(q)
}
