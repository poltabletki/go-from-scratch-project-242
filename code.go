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
		return formatSize(info.Size(), human), nil
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

	return formatSize(totalSize, human), nil
}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	value := float64(size)
	unitIndex := 0

	for value >= 1024 && unitIndex < len(units)-1 {
		value /= 1024
		unitIndex++
	}

	if unitIndex == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", value, units[unitIndex])
}
