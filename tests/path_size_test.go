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

func TestGetPathSize_HiddenFilesFiltering(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()

	visibleFile := filepath.Join(tempDir, "visible.txt")
	require.NoError(t, os.WriteFile(visibleFile, []byte("12345"), 0o644))

	hiddenFile := filepath.Join(tempDir, ".hidden.txt")
	require.NoError(t, os.WriteFile(hiddenFile, []byte("123"), 0o644))

	hiddenDir := filepath.Join(tempDir, ".hidden-dir")
	require.NoError(t, os.Mkdir(hiddenDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(hiddenDir, "inside.txt"), []byte("123456"), 0o644))

	sizeWithoutHidden, err := code.GetPathSize(tempDir, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "5B", sizeWithoutHidden)

	sizeWithHidden, err := code.GetPathSize(tempDir, false, false, true)
	require.NoError(t, err)
	require.Equal(t, "8B", sizeWithHidden)
}

func TestGetPathSize_DirectoryRecursive(t *testing.T) {
	t.Parallel()

	path := filepath.Join("..", "testdata", "fixtures", "dir-a")
	size, err := code.GetPathSize(path, true, false, false)

	require.NoError(t, err)
	require.Equal(t, "57B", size)
}

func TestGetPathSize_RecursiveHiddenDirectoriesFiltering(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()

	require.NoError(t, os.WriteFile(filepath.Join(tempDir, "visible.txt"), []byte("12345"), 0o644))

	hiddenDir := filepath.Join(tempDir, ".hidden-dir")
	require.NoError(t, os.Mkdir(hiddenDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(hiddenDir, "inside.txt"), []byte("123456"), 0o644))

	sizeWithoutHidden, err := code.GetPathSize(tempDir, true, false, false)
	require.NoError(t, err)
	require.Equal(t, "5B", sizeWithoutHidden)

	sizeWithHidden, err := code.GetPathSize(tempDir, true, false, true)
	require.NoError(t, err)
	require.Equal(t, "11B", sizeWithHidden)
}
