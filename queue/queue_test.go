package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	q := Queue[int]{}
	require.Equal(t, 0, q.Count())

	q.Enqueue(1)
	require.Equal(t, []int{1}, q.items)
	require.Equal(t, 1, q.Count())

	q.Enqueue(2)
	require.Equal(t, []int{1, 2}, q.items)
	require.Equal(t, 2, q.Count())

	require.Equal(t, 1, q.Dequeue())
	require.Equal(t, []int{2}, q.items)
	require.Equal(t, 1, q.Count())

	require.Equal(t, 2, q.Dequeue())
	require.Equal(t, []int{}, q.items)
	require.Equal(t, 0, q.Count())

	require.Panics(t, func() { q.Dequeue() }, "qeueue is empty")
}
