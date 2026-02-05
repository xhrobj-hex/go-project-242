package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize вычисляет размер файла или директории и возвращает его в виде строки.
//
// Функция поддерживает:
//   - рекурсивный обход директорий, если recursive == true
//   - учёт скрытых файлов и директорий (имена, начинающиеся с '.'), если all == true
//   - человекочитаемый формат размера (KB, MB, GB и т.д.), если human == true
//
// Если path указывает на файл, возвращается его размер.
// Если path указывает на директорию, вычисляется размер её содержимого.
// При recursive == false учитываются только файлы первого уровня.
//
// Возвращаемое значение содержит только отформатированный размер
// (например: "42B", "1.5MB").
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
}

func getSize(path string, recursive, includeHidden bool) (int64, error) {
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
				subDirSize, err := getSize(subPath, recursive, includeHidden)
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

func formatSize(size int64, human bool) string {
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
