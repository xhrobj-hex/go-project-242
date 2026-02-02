# Анализатор размера диска (Go)

CLI-утилита для подсчёта размера файлов и директорий с поддержкой:
- человекочитаемого формата
- скрытых файлов
- рекурсивного обхода директорий

## Возможности

- 📁 Подсчёт размера файла или директории
- 🔁 Рекурсивный подсчёт вложенных директорий
- 👁 Поддержка скрытых файлов (`.`)
- 📊 Человекочитаемый формат (`KB`, `MB`, `GB`, ...)
- ✅ Покрытие тестами
- 🤖 CI через GitHub Actions

## Установка

### Сборка из исходников

```bash
git clone https://github.com/xhrobj-hex/go-project-242.git
cd go-project-242
make build
```

Бинарник появится в `./bin/hexlet-path-size`.


## Использование

```bash
./bin/hexlet-path-size <path> [flags]
```

### Примеры

Подсчёт размера файла:

```bash
./bin/hexlet-path-size file.txt
```

```
123B	file.txt
```

Человекочитаемый формат:

```bash
./bin/hexlet-path-size --human file.txt
```

```
1.2KB	file.txt
```

Подсчёт директории без скрытых файлов:

```bash
./bin/hexlet-path-size project/
```

Подсчёт с учётом скрытых файлов:

```bash
./bin/hexlet-path-size project/ -a
```

Рекурсивный подсчёт:

```bash
./bin/hexlet-path-size project/ -r
```

Комбинация флагов:

```bash
./bin/hexlet-path-size project/ -H -a -r
```

## Флаги

| Флаг | Алиас | Описание |
|-----|------|----------|
| `--human` | `-H` | Человекочитаемый формат размеров |
| `--all` | `-a` | Учитывать скрытые файлы и директории |
| `--recursive` | `-r` | Рекурсивный подсчёт директорий |
| `--help` | `-h` | Справка |

## Запуск тестов

```bash
make test
```

## Демонстрация

🎬 Видео с примером установки и работы программы записано с помощью **asciinema**:

[![asciinema demo](https://asciinema.org/a/WlDIau0aBxDQxXdS.svg)](https://asciinema.org/a/WlDIau0aBxDQxXdS)

---

### Hexlet tests and linter status:
[![Actions Status](https://github.com/xhrobj-hex/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/xhrobj-hex/go-project-242/actions)

### Project CI - lint & tests
[![(-_-) go-ci](https://github.com/xhrobj-hex/go-project-242/actions/workflows/go-ci.yml/badge.svg)](https://github.com/xhrobj-hex/go-project-242/actions/workflows/go-ci.yml)
