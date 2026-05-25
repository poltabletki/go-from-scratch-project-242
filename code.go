package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return fmt.Sprintf("%dB", info.Size()), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var totalSize int64
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		entryInfo, err := entry.Info()
		if err != nil {
			return "", err
		}

		totalSize += entryInfo.Size()
	}

	return fmt.Sprintf("%dB", totalSize), nil
}
