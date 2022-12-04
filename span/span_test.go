package span

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpanContains(t *testing.T) {
	for _, data := range []struct {
		From, To Span
		expected bool
	}{
		{Span{1, 3}, Span{1, 3}, true},
		{Span{1, 4}, Span{2, 3}, true},
		{Span{1, 3}, Span{2, 2}, true},
		{Span{1, 3}, Span{4, 6}, false},
	} {
		require.Equal(t, data.From.Contains(data.To), data.expected)
		require.Equal(t, data.To.Contains(data.From), data.expected)
	}
}

func TestSpanOverlaps(t *testing.T) {
	for _, data := range []struct {
		From, To Span
		expected bool
	}{
		{Span{1, 3}, Span{1, 3}, true},
		{Span{1, 3}, Span{3, 5}, true},
		{Span{1, 3}, Span{2, 3}, true},
		{Span{1, 6}, Span{2, 3}, true},
	} {
		require.Equal(t, data.From.Overlaps(data.To), data.expected)
		require.Equal(t, data.To.Overlaps(data.From), data.expected)
	}
}
