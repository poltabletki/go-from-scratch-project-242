package tests

import (
	"os"
	"path/filepath"
	"testing"

	"code"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	t.Parallel()

	path := filepath.Join("..", "testdata", "fixtures", "file-root.txt")
	size, err := code.GetPathSize(path, false, false, false)

	require.NoError(t, err)
	require.Equal(t, "5B", size)
}

func TestGetPathSize_Directory(t *testing.T) {
	t.Parallel()

	path := filepath.Join("..", "testdata", "fixtures", "dir-a")
	size, err := code.GetPathSize(path, false, false, false)

	require.NoError(t, err)
	// alpha.txt (4B) + beta.txt (6B), nested directory is ignored in non-recursive mode.
	require.Equal(t, "10B", size)
}

func TestGetPathSize_FileHumanReadable(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "big.bin")

	f, err := os.Create(filePath)
	require.NoError(t, err)
	require.NoError(t, f.Truncate(1234567))
	require.NoError(t, f.Close())

	size, err := code.GetPathSize(filePath, false, true, false)
	require.NoError(t, err)
	require.Equal(t, "1.2MB", size)
}

func TestGetPathSize_DirectoryHumanReadable(t *testing.T) {
	t.Parallel()

	path := filepath.Join("..", "testdata", "fixtures", "dir-a")
	size, err := code.GetPathSize(path, false, true, false)

	require.NoError(t, err)
	require.Equal(t, "10B", size)
}

func TestGetPathSize_NotFound(t *testing.T) {
	t.Parallel()

	path := filepath.Join("..", "testdata", "fixtures", "missing.txt")
	size, err := code.GetPathSize(path, false, false, false)

	require.Error(t, err)
	require.Empty(t, size)
}
