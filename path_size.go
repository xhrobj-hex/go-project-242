package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string, human bool) (string, error) {
	size, err := GetSize(path)
	if err != nil {
		return "", err
	}

	r := formatPathSize(size, path)
	return r, nil
}

func GetSize(path string) (int64, error) {
	pathStat, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if !pathStat.IsDir() {
		r := pathStat.Size()
		return r, nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var size int64
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		entryInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}
		size += entryInfo.Size()
	}

	return size, nil
}

func FormatSize(size int64, human bool) string {
	return fmt.Sprintf("%dB", size)
}

func formatPathSize(size int64, path string) string {
	return fmt.Sprintf("%d\t%s", size, path)
}
