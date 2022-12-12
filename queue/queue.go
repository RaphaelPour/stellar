package queue

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T]{
	q := new(Queue[T])
	q.items = make([]T,0)
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.items) == 0 {
		panic("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q Queue[T]) Count() int {
	return len(q.items)
}
