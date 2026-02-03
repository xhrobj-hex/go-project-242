package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return FormatSize(size, human), nil
}

func GetSize(path string, recursive, includeHidden bool) (int64, error) {
	const hiddenPrefix = "."

	pathStat, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !includeHidden && strings.HasPrefix(pathStat.Name(), hiddenPrefix) {
		return 0, nil
	}

	if !pathStat.IsDir() {
		return pathStat.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var size int64
	for _, entry := range entries {
		if !includeHidden && strings.HasPrefix(entry.Name(), hiddenPrefix) {
			continue
		}
		if entry.IsDir() {
			if recursive {
				subPath := filepath.Join(path, entry.Name())
				subDirSize, err := GetSize(subPath, recursive, includeHidden)
				if err != nil {
					return 0, err
				}
				size += subDirSize
			}
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
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	const (
		KB int64 = 1024
		MB       = KB * 1024
		GB       = MB * 1024
		TB       = GB * 1024
		PB       = TB * 1024
		EB       = PB * 1024
	)

	type unit struct {
		name  string
		value int64
	}

	units := []unit{
		{"EB", EB},
		{"PB", PB},
		{"TB", TB},
		{"GB", GB},
		{"MB", MB},
		{"KB", KB},
	}

	for _, u := range units {
		if size < u.value {
			continue
		}
		x := float64(size) / float64(u.value)
		return fmt.Sprintf("%.1f%s", x, u.name)
	}

	return fmt.Sprintf("%dB", size)
}
