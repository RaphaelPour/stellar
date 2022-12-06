package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStackPop(t *testing.T) {
	s := Stack[string]{
		items: []string{"a"},
	}
	require.Equal(t, "a", s.Pop())
}

func TestStackPopPanics(t *testing.T) {
	s := Stack[string]{}
	require.Panics(t, func() { s.Pop() }, "stack is empty")
}

func TestStackPopN(t *testing.T) {
	s := Stack[string]{
		items: []string{"a", "b", "c"},
	}
	require.Equal(t, []string{"a", "b", "c"}, s.PopN(3))
}

func TestStackPopNPanics(t *testing.T) {
	s := Stack[string]{}
	require.Panics(t, func() { s.PopN(100) }, "stack has not enough items")
}

func TestStackPush(t *testing.T) {
	s := Stack[string]{}
	s.Push("a")
	require.Equal(t, []string{"a"}, s.items)
}

func TestStackPushAhead(t *testing.T) {
	s := Stack[string]{}
	s.PushAhead("c")
	s.PushAhead("b")
	s.PushAhead("a")
	require.Equal(t, []string{"a", "b", "c"}, s.items)
}

func TestStackPushN(t *testing.T) {
	s := Stack[string]{}
	s.PushN([]string{"a", "b", "c"})
	require.Equal(t, []string{"a", "b", "c"}, s.items)
}

func TestStackClear(t *testing.T) {
	s := Stack[string]{
		items: []string{"something"},
	}

	require.Equal(t, []string{"something"}, s.items)

	s.Clear()

	require.Empty(t, s.items)
}
