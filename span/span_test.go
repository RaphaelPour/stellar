package span

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpanContains(t *testing.T) {
	for _, data := range []struct {
		from, to Span
		expected bool
	}{
		{Span{1, 3}, Span{1, 3}, true},
		{Span{1, 4}, Span{2, 3}, true},
		{Span{1, 3}, Span{2, 2}, true},
		{Span{1, 3}, Span{4, 6}, false},
	} {
		require.Equal(t, data.from.Contains(data.to), data.expected)
		require.Equal(t, data.to.Contains(data.from), data.expected)
	}
}

func TestSpanOverlaps(t *testing.T) {
	for _, data := range []struct {
		from, to Span
		expected bool
	}{
		{Span{1, 3}, Span{1, 3}, true},
		{Span{1, 3}, Span{3, 5}, true},
		{Span{1, 3}, Span{2, 3}, true},
		{Span{1, 6}, Span{2, 3}, true},
	} {
		require.Equal(t, data.from.Overlaps(data.to), data.expected)
		require.Equal(t, data.to.Overlaps(data.from), data.expected)
	}
}
