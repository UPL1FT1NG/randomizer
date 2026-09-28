# randomizer - Build Instructions

## Requirements
- Go 1.21 or higher
- Windows 10/11
- Git (for downloading dependencies)

## Build Steps

### 1. Initialize Project
```cmd
go mod download
```

### 2. Build Executable (without console window)
```cmd
go build -ldflags="-H windowsgui -s -w" -o randomizer.exe
```

**Compilation flags:**
- `-H windowsgui` — disables console window (GUI mode)
- `-s` — removes symbol table (reduces size)
- `-w` — removes DWARF debug information (reduces size)

### 3. Run
```cmd
randomizer.exe
```

## Application Controls

### Hotkeys and Mouse:
- **LMB (hold and drag)** — move window across screen
- **RMB (click)** — open context menu (Interval, Color, Exit)
- **RMB (hold and move left/right)** — resize window (35-500 px)
- **MMB / ESC** — close application

### Features:
- Window always on top (topmost)
- New number generation every 5 seconds (default)
- Color changes by gradient: Green → Lime → Yellow → Orange → Red
- Font size automatically scales to fill 92% of window area
- Square window shape (1:1) without borders
- Native Win32 context menu (zero footprint, RAM-only settings)

## Alternative Build (with console for debugging)
```cmd
go build -o randomizer-debug.exe
```

## Binary Size
Expected size: ~2 MB (thanks to pure Win32 API without CGO)

---

# randomizer - Инструкции по сборке

## Требования
- Go 1.21 или выше
- Windows 10/11
- Git (для загрузки зависимостей)

## Шаги сборки

### 1. Инициализация проекта
```cmd
go mod download
```

### 2. Сборка исполняемого файла (без консольного окна)
```cmd
go build -ldflags="-H windowsgui -s -w" -o randomizer.exe
```

**Параметры компиляции:**
- `-H windowsgui` — отключает консольное окно (GUI-режим)
- `-s` — удаляет таблицу символов (уменьшает размер)
- `-w` — удаляет отладочную информацию DWARF (уменьшает размер)

### 3. Запуск
```cmd
randomizer.exe
```

## Управление приложением

### Горячие клавиши и мышь:
- **ЛКМ (зажать и тянуть)** — перемещение окна по экрану
- **ПКМ (клик)** — открыть контекстное меню (Interval, Color, Exit)
- **ПКМ (зажать и двигать влево/вправо)** — изменение размера окна (35-500 px)
- **СКМ / ESC** — закрытие приложения

### Особенности:
- Окно всегда поверх других окон (topmost)
- Генерация нового числа каждые 5 секунд (по умолчанию)
- Цвет числа меняется по градиенту: Зеленый → Салатовый → Желтый → Оранжевый → Красный
- Размер шрифта автоматически масштабируется для заполнения 92% площади окна
- Квадратная форма окна (1:1) без рамок
- Нативное Win32 контекстное меню (нулевой след, настройки только в RAM)

## Альтернативная сборка (с консолью для отладки)
```cmd
go build -o randomizer-debug.exe
```

## Размер бинарника
Ожидаемый размер: ~2 МБ (благодаря использованию чистого Win32 API без CGO)
