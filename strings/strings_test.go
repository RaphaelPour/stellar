package strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T) {
	require.Equal(t, "cba", Reverse("abc"))
	require.Equal(t, "ba", Reverse("ab"))
	require.Equal(t, "a", Reverse("a"))
}

func TestToInt(t *testing.T) {
	require.Equal(t, 10, ToInt("10"))
}
