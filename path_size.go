package code

/*

Задачи
------
Добавьте функцию GetSize(), которая принимает путь к файлу или директории и возвращает размер:

Если путь — это файл, вернуть его размер.
Если директория — суммировать размеры файлов первого уровня.
Измените вывод программы, чтобы результат выводился в формате:

<размер>  <путь>
В качестве разделителя между размером и путем используется символ табуляции.

Подсказки
---------
os.Lstat(path), чтобы получить FileInfo.
entry.Info() для получения размера.

*/

import (
	"fmt"
	"os"
)

func GetSize(path string) (string, error) {
	pathStat, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	if !pathStat.IsDir() {
		r := formatResult(pathStat.Size(), path)
		return r, nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var size int64
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		entryInfo, err := entry.Info()
		if err != nil {
			return "", err
		}
		size += entryInfo.Size()
	}

	r := formatResult(size, path)
	return r, nil
}

func formatResult(size int64, path string) string {
	return fmt.Sprintf("%d\t%s", size, path)
}
