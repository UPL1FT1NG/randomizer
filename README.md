<div align="center">

# Randomizer

### A lightweight frameless overlay randomizer for Windows built with pure Go and Win32 API

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/) [![Platform](https://img.shields.io/badge/Platform-Windows-blue?style=flat&logo=windows)](https://www.microsoft.com/windows) [![Release](https://img.shields.io/github/v/release/UPL1FT1NG/randomizer?color=orange&style=flat)](https://github.com/UPL1FT1NG/randomizer/releases/latest) [![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

<strong>English</strong> &nbsp;•&nbsp; <a href="#russian">🇷🇺 Читать на русском</a>

| Dynamic Gradient | Context Menu |
| :---: | :---: |
| <img src="https://github.com/user-attachments/assets/77f8d4a5-f089-406f-84ba-45b41e819829" width="154" alt="Low number 24" /> <img src="https://github.com/user-attachments/assets/f522dc9f-bc01-4509-b3d9-aa88ed93f83c" width="146" alt="Mid number 53" /> <img src="https://github.com/user-attachments/assets/c3e9e4c8-c229-4809-8fe9-9350561760f6" width="153" alt="High number 97" /> | <img src="https://github.com/user-attachments/assets/ea22672c-6604-441c-b384-554cac9466f5" width="152" alt="Interval menu" /> <img src="https://github.com/user-attachments/assets/de4be81c-2572-42bd-9cc4-75f8b8a35a14" width="150" alt="Color menu" /> |

<br/>

<a href="https://github.com/UPL1FT1NG/randomizer/releases/latest/download/randomizer.exe"><img src="https://img.shields.io/badge/Download-randomizer.exe-0078D4?style=for-the-badge&logo=windows&logoColor=white" alt="Download randomizer.exe" /></a>&nbsp;&nbsp;<a href="https://github.com/UPL1FT1NG/randomizer/releases/latest"><img src="https://img.shields.io/github/v/release/UPL1FT1NG/randomizer?label=Releases&style=for-the-badge&logo=github&logoColor=white&color=24292e" alt="View Releases" /></a>

</div>

---

## Description

Ultra-minimalist overlay application that generates random numbers from 1 to 100 with configurable intervals. The window always stays on top of other applications and is hidden from the taskbar.

### Features

- **Always-on-Top & Multi-Instance** — frameless overlay widget; run multiple independent windows simultaneously
- **Dynamic Gradient & Auto-Fit** — numbers smoothly transition from green to red (invertible); font dynamically scales to fill window
- **Smooth Resizing & Drag** — hold LMB to move anywhere, hold RMB to smoothly resize from 35×35 to 500×500 px
- **Instant Controls (Menu & CLI)** — right-click to switch intervals (1s–10s) and palette, or preconfigure via command-line flags
- **Zero Footprint & Portable** — standalone ~2 MB binary with no installation, zero disk/registry writes, and flicker-free rendering

## Controls

| Action | Description |
|--------|-------------|
| **Left Mouse Button (LMB)** | Hold and drag to move window across screen |
| **Right Mouse Button (RMB Click)** | Open context menu (Interval, Color, Exit) |
| **Right Mouse Button (RMB Drag)** | Hold and move left/right to resize window (35-500 px) |
| **Middle Mouse Button (MMB)** | Instant exit |
| **ESC** | Instant exit |

### Context Menu

Right-click on the window to access:
- **Interval** — Change update frequency (1s, 3s, 5s, 7s, 10s)
- **Color** — Switch color scheme:
  - Green → Red — default
  - Red → Green — inverted
- **Exit** — Close application

## CLI Flags / Shortcuts

Launch with custom settings:

```cmd
# Default: 5-second interval, Green → Red palette
randomizer.exe

# Custom interval (7 seconds)
randomizer.exe -t 7

# Inverted color palette (Red → Green)
randomizer.exe -inv

# Combine both options
randomizer.exe -t 10 -inv
```

**Options:**
- `-t <int>` — Update interval in seconds (default: 5, range: 1-60)
- `-inv` — Invert color palette (1=Red, 100=Green)

## Build

### Requirements
- Go 1.21 or higher
- Windows 10/11

### Commands

```cmd
# Download dependencies
go mod download

# Build executable (without console window)
go build -ldflags="-H windowsgui -s -w" -o randomizer.exe
```

**Compilation flags:**
- `-H windowsgui` — disables console window (GUI mode)
- `-s` — removes symbol table
- `-w` — removes DWARF debug information

## Technical Details

- **Language:** Go 1.21+
- **API:** Pure Win32 API via `golang.org/x/sys/windows`
- **Font:** Arial Black with FW_BOLD
- **Quality:** CLEARTYPE_QUALITY for maximum sharpness
- **Initial size:** 100×100 px (screen center)
- **Default interval:** 5 seconds
- **Timer:** Win32 SetTimer (configurable: 1s, 3s, 5s, 7s, 10s)
- **Rendering:** Double buffering via CreateCompatibleDC
- **Auto-fit algorithm:** Dynamic font sizing using GetTextExtentPoint32W
- **Context menu:** CreatePopupMenu with TrackPopupMenuEx (destroyed after use)
- **Stability:** runtime.LockOSThread(), no GDI leaks, smooth gradient

## License

MIT

---

<a id="russian"></a>

<div align="center">

# Randomizer

### Легковесный frameless оверлей-рандомайзер для Windows на чистом Go с Win32 API

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/) [![Platform](https://img.shields.io/badge/Platform-Windows-blue?style=flat&logo=windows)](https://www.microsoft.com/windows) [![Release](https://img.shields.io/github/v/release/UPL1FT1NG/randomizer?color=orange&style=flat)](https://github.com/UPL1FT1NG/randomizer/releases/latest) [![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

<a href="#randomizer">🇬🇧 Back to English</a> &nbsp;•&nbsp; <strong>🇷🇺 Русский</strong>

| Dynamic Gradient | Context Menu |
| :---: | :---: |
| <img src="https://github.com/user-attachments/assets/77f8d4a5-f089-406f-84ba-45b41e819829" width="154" alt="Малое число 24" /> <img src="https://github.com/user-attachments/assets/f522dc9f-bc01-4509-b3d9-aa88ed93f83c" width="146" alt="Среднее число 53" /> <img src="https://github.com/user-attachments/assets/c3e9e4c8-c229-4809-8fe9-9350561760f6" width="153" alt="Большое число 97" /> | <img src="https://github.com/user-attachments/assets/ea22672c-6604-441c-b384-554cac9466f5" width="152" alt="Меню интервалов" /> <img src="https://github.com/user-attachments/assets/de4be81c-2572-42bd-9cc4-75f8b8a35a14" width="150" alt="Меню цветов" /> |

<br/>

<a href="https://github.com/UPL1FT1NG/randomizer/releases/latest/download/randomizer.exe"><img src="https://img.shields.io/badge/Скачать-randomizer.exe-0078D4?style=for-the-badge&logo=windows&logoColor=white" alt="Скачать randomizer.exe" /></a>&nbsp;&nbsp;<a href="https://github.com/UPL1FT1NG/randomizer/releases/latest"><img src="https://img.shields.io/github/v/release/UPL1FT1NG/randomizer?label=%D0%A0%D0%B5%D0%BB%D0%B8%D0%B7%D1%8B&style=for-the-badge&logo=github&logoColor=white&color=24292e" alt="Посмотреть релизы" /></a>

</div>

---

## Описание

Ультраминималистичное приложение-оверлей, генерирующее случайные числа от 1 до 100 с настраиваемым интервалом. Окно всегда находится поверх других приложений и скрыто с панели задач.

### Особенности

- **Поверх всех окон и Multi-Instance** — компактный безрамочный оверлей с поддержкой запуска нескольких независимых окон
- **Динамический градиент и Auto-Fit** — плавный переход цвета от зеленого к красному (с инверсией); автоподгонка размера шрифта
- **Плавный ресайз и перемещение** — перемещение зажатием ЛКМ, плавное изменение размера от 35×35 до 500×500 px зажатием ПКМ
- **Быстрое управление (Меню и CLI)** — нативное контекстное меню (интервалы 1–10с, темы) и поддержка флагов командной строки
- **Zero Footprint и портативность** — автономный бинарник ~2 МБ без установки, без записей в реестр и без мерцания

## Управление

| Действие | Описание |
|----------|----------|
| **Левая кнопка мыши (ЛКМ)** | Зажать и тянуть для перемещения окна по экрану |
| **Правая кнопка мыши (ПКМ клик)** | Открыть контекстное меню (Interval, Color, Exit) |
| **Правая кнопка мыши (ПКМ драг)** | Зажать и двигать влево/вправо для изменения размера окна (35-500 px) |
| **Средняя кнопка мыши (СКМ)** | Мгновенный выход |
| **ESC** | Мгновенный выход |

### Контекстное меню

Кликните правой кнопкой мыши по окну для доступа к:
- **Interval** — Изменение частоты обновления (1s, 3s, 5s, 7s, 10s)
- **Color** — Переключение цветовой схемы:
  - Green → Red — по умолчанию
  - Red → Green — инвертированная
- **Exit** — Закрытие приложения

## CLI Флаги / Ярлыки

Запуск с пользовательскими настройками:

```cmd
# По умолчанию: интервал 5 секунд, палитра Green → Red
randomizer.exe

# Пользовательский интервал (7 секунд)
randomizer.exe -t 7

# Инвертированная цветовая палитра (Red → Green)
randomizer.exe -inv

# Комбинация обеих опций
randomizer.exe -t 10 -inv
```

**Опции:**
- `-t <int>` — Интервал обновления в секундах (по умолчанию: 5, диапазон: 1-60)
- `-inv` — Инвертировать цветовую палитру (1=Красный, 100=Зеленый)

## Сборка

### Требования
- Go 1.21 или выше
- Windows 10/11

### Команды

```cmd
# Скачать зависимости
go mod download

# Собрать исполняемый файл (без консольного окна)
go build -ldflags="-H windowsgui -s -w" -o randomizer.exe
```

**Параметры компиляции:**
- `-H windowsgui` — отключает консольное окно (GUI-режим)
- `-s` — удаляет таблицу символов
- `-w` — удаляет отладочную информацию DWARF

## Технические детали

- **Язык:** Go 1.21+
- **API:** Чистый Win32 API через `golang.org/x/sys/windows`
- **Шрифт:** Arial Black с FW_BOLD
- **Качество:** CLEARTYPE_QUALITY для максимальной четкости
- **Стартовый размер:** 100×100 px (центр экрана)
- **Интервал по умолчанию:** 5 секунд
- **Таймер:** Win32 SetTimer (настраиваемый: 1s, 3s, 5s, 7s, 10s)
- **Отрисовка:** Двойная буферизация через CreateCompatibleDC
- **Алгоритм Auto-fit:** Динамическое масштабирование шрифта через GetTextExtentPoint32W
- **Контекстное меню:** CreatePopupMenu с TrackPopupMenuEx (уничтожается после использования)
- **Стабильность:** runtime.LockOSThread(), отсутствие утечек GDI, плавный градиент

## Лицензия

MIT
