package hack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBool(t *testing.T) {
	require.True(t, Bool(int(1)))
	require.False(t, Bool(int(0)))
	require.True(t, Bool(int8(1)))
	require.False(t, Bool(int8(0)))
	require.True(t, Bool(int16(1)))
	require.False(t, Bool(int16(0)))
	require.True(t, Bool(int32(1)))
	require.False(t, Bool(int32(0)))
	require.True(t, Bool(int64(1)))
	require.False(t, Bool(int64(0)))
	require.True(t, Bool(uint(1)))
	require.False(t, Bool(uint(0)))
	require.True(t, Bool(uint8(1)))
	require.False(t, Bool(uint8(0)))
	require.True(t, Bool(uint16(1)))
	require.False(t, Bool(uint16(0)))
	require.True(t, Bool(uint32(1)))
	require.False(t, Bool(uint32(0)))
	require.True(t, Bool(uint64(1)))
	require.False(t, Bool(uint64(0)))
	require.True(t, Bool(float32(1)))
	require.False(t, Bool(float32(0)))
	require.True(t, Bool(float64(1)))
	require.False(t, Bool(float64(0)))
}

func TestWormhole(t *testing.T) {
	require.Equal(t, int(0), Wormhole(false))
	require.Equal(t, int(1), Wormhole(true))
}
