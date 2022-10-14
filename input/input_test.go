package input

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "input-")
	require.Nil(t, err)
	defer os.Remove(tmpFile.Name())

	input := "test"
	_, err = tmpFile.Write([]byte(input))
	require.Nil(t, err)
	require.Nil(t, tmpFile.Close())

	require.Equal(t, []string{input}, LoadString(tmpFile.Name()))
}

func TestLoadDefault(t *testing.T) {
	file, err := os.Create("input")
	require.Nil(t, err)
	defer os.Remove("input")

	input := "test"
	_, err = file.Write([]byte(input))
	require.Nil(t, err)
	require.Nil(t, file.Close())

	require.Equal(t, []string{input}, LoadDefaultString())
}

func TestLoadInt(t *testing.T) {
	file, err := os.Create("input")
	require.Nil(t, err)
	defer os.Remove("input")

	input := "1\n2"
	_, err = file.Write([]byte(input))
	require.Nil(t, err)
	require.Nil(t, file.Close())

	require.Equal(t, []int{1, 2}, LoadDefaultInt())
}

func TestLoadIntList(t *testing.T) {
	file, err := os.Create("input")
	require.Nil(t, err)
	defer os.Remove("input")

	input := "1,2"
	_, err = file.Write([]byte(input))
	require.Nil(t, err)
	require.Nil(t, file.Close())

	require.Equal(t, []int{1, 2}, LoadIntList("input"))
}
