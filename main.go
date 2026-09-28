// randomizer v1.0 - minimalist overlay number randomizer for Windows
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Win32 API constants
const (
	WS_POPUP            = 0x80000000
	WS_VISIBLE          = 0x10000000
	WS_EX_TOPMOST       = 0x00000008
	WS_EX_TOOLWINDOW    = 0x00000080
	WM_DESTROY          = 0x0002
	WM_PAINT            = 0x000F
	WM_TIMER            = 0x0113
	WM_ERASEBKGND       = 0x0014
	WM_LBUTTONDOWN      = 0x0201
	WM_RBUTTONDOWN      = 0x0204
	WM_RBUTTONUP        = 0x0205
	WM_MBUTTONUP        = 0x0208
	WM_MOUSEMOVE        = 0x0200
	WM_KEYDOWN          = 0x0100
	WM_COMMAND          = 0x0111
	WM_NCLBUTTONDOWN    = 0x00A1
	WM_SETCURSOR        = 0x0020
	VK_ESCAPE           = 0x1B
	HTCAPTION           = 2
	IDC_ARROW           = 32512
	DT_CENTER           = 0x00000001
	DT_VCENTER          = 0x00000004
	DT_SINGLELINE       = 0x00000020
	DT_NOPREFIX         = 0x00000800
	FW_BOLD             = 700
	FW_BLACK            = 900
	TRANSPARENT         = 1
	DEFAULT_CHARSET     = 1
	OUT_TT_PRECIS       = 4
	CLIP_DEFAULT_PRECIS = 0
	ANTIALIASED_QUALITY = 4
	CLEARTYPE_QUALITY   = 5
	VARIABLE_PITCH      = 2
	FF_SWISS            = 32
	SM_CXSCREEN         = 0
	SM_CYSCREEN         = 1
	SW_SHOW             = 5
	SWP_NOMOVE          = 0x0002
	SWP_NOZORDER        = 0x0004
	SRCCOPY             = 0x00CC0020
	TPM_RIGHTBUTTON     = 0x0002
	TPM_TOPALIGN        = 0x0000
	TPM_RETURNCMD       = 0x0100
	MF_STRING           = 0x0000
	MF_SEPARATOR        = 0x0800
	MF_CHECKED          = 0x0008
	MF_POPUP            = 0x0010
	TIMER_ID            = 1
)

// Menu command IDs
const (
	CMD_INTERVAL_1S  = 1001
	CMD_INTERVAL_3S  = 1003
	CMD_INTERVAL_5S  = 1005
	CMD_INTERVAL_7S  = 1007
	CMD_INTERVAL_10S = 1010
	CMD_PALETTE_NORM = 2001
	CMD_PALETTE_INV  = 2002
	CMD_EXIT         = 3001
)

var (
	kernel32                   = windows.NewLazySystemDLL("kernel32.dll")
	user32                     = windows.NewLazySystemDLL("user32.dll")
	gdi32                      = windows.NewLazySystemDLL("gdi32.dll")
	procGetModuleHandleW       = kernel32.NewProc("GetModuleHandleW")
	procRegisterClassExW       = user32.NewProc("RegisterClassExW")
	procCreateWindowExW        = user32.NewProc("CreateWindowExW")
	procDefWindowProcW         = user32.NewProc("DefWindowProcW")
	procGetMessageW            = user32.NewProc("GetMessageW")
	procTranslateMessage       = user32.NewProc("TranslateMessage")
	procDispatchMessageW       = user32.NewProc("DispatchMessageW")
	procPostQuitMessage        = user32.NewProc("PostQuitMessage")
	procBeginPaint             = user32.NewProc("BeginPaint")
	procEndPaint               = user32.NewProc("EndPaint")
	procInvalidateRect         = user32.NewProc("InvalidateRect")
	procUpdateWindow           = user32.NewProc("UpdateWindow")
	procGetClientRect          = user32.NewProc("GetClientRect")
	procSetCapture             = user32.NewProc("SetCapture")
	procReleaseCapture         = user32.NewProc("ReleaseCapture")
	procGetCursorPos           = user32.NewProc("GetCursorPos")
	procGetWindowRect          = user32.NewProc("GetWindowRect")
	procSetWindowPos           = user32.NewProc("SetWindowPos")
	procShowWindow             = user32.NewProc("ShowWindow")
	procSetForegroundWindow    = user32.NewProc("SetForegroundWindow")
	procGetSystemMetrics       = user32.NewProc("GetSystemMetrics")
	procSendMessageW           = user32.NewProc("SendMessageW")
	procSetTimer               = user32.NewProc("SetTimer")
	procKillTimer              = user32.NewProc("KillTimer")
	procLoadCursorW            = user32.NewProc("LoadCursorW")
	procSetCursor              = user32.NewProc("SetCursor")
	procCreatePopupMenu        = user32.NewProc("CreatePopupMenu")
	procAppendMenuW            = user32.NewProc("AppendMenuW")
	procTrackPopupMenuEx       = user32.NewProc("TrackPopupMenuEx")
	procDestroyMenu            = user32.NewProc("DestroyMenu")
	procDestroyWindow          = user32.NewProc("DestroyWindow")
	procCreateSolidBrush       = gdi32.NewProc("CreateSolidBrush")
	procCreateCompatibleDC     = gdi32.NewProc("CreateCompatibleDC")
	procCreateCompatibleBitmap = gdi32.NewProc("CreateCompatibleBitmap")
	procSelectObject           = gdi32.NewProc("SelectObject")
	procBitBlt                 = gdi32.NewProc("BitBlt")
	procDeleteDC               = gdi32.NewProc("DeleteDC")
	procDeleteObject           = gdi32.NewProc("DeleteObject")
	procSetTextColor           = gdi32.NewProc("SetTextColor")
	procSetBkMode              = gdi32.NewProc("SetBkMode")
	procCreateFontW            = gdi32.NewProc("CreateFontW")
	procGetTextExtentPoint32W  = gdi32.NewProc("GetTextExtentPoint32W")
	procDrawTextW              = user32.NewProc("DrawTextW")
	procFillRect               = user32.NewProc("FillRect")
)

// WNDCLASSEXW structure
type WNDCLASSEXW struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     windows.Handle
	HIcon         windows.Handle
	HCursor       windows.Handle
	HbrBackground windows.Handle
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       windows.Handle
}

// MSG structure
type MSG struct {
	Hwnd    windows.Handle
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

// POINT structure
type POINT struct {
	X, Y int32
}

// RECT structure
type RECT struct {
	Left, Top, Right, Bottom int32
}

// PAINTSTRUCT structure
type PAINTSTRUCT struct {
	Hdc         windows.Handle
	FErase      int32
	RcPaint     RECT
	FRestore    int32
	FIncUpdate  int32
	RgbReserved [32]byte
}

// SIZE structure
type SIZE struct {
	CX, CY int32
}

// Global state
var (
	currentNumber      int   = 50
	windowSize         int32 = 100
	isRDown            bool  = false
	isDragging         bool  = false
	resizeStartX       int32 = 0
	resizeStartY       int32 = 0
	resizeStartSize    int32 = 0
	hwnd               windows.Handle
	currentIntervalSec uint32 = 5
	isInverted         bool   = false
)

func main() {
	// CRITICAL: Lock OS thread for Win32 message loop
	runtime.LockOSThread()

	// Parse CLI flags
	intervalFlag := flag.Int("t", 5, "Update interval in seconds (1-60)")
	invertFlag := flag.Bool("inv", false, "Invert color palette (1=Red, 100=Green)")
	flag.Parse()

	// Validate and set interval
	if *intervalFlag < 1 {
		currentIntervalSec = 1
	} else if *intervalFlag > 60 {
		currentIntervalSec = 60
	} else {
		currentIntervalSec = uint32(*intervalFlag)
	}

	// Set inversion flag
	isInverted = *invertFlag

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate first number immediately
	currentNumber = rand.Intn(100) + 1

	// Get module handle
	hInstance, _, _ := procGetModuleHandleW.Call(0)

	// Load cursor
	hCursor, _, _ := procLoadCursorW.Call(0, IDC_ARROW)

	// Register window class
	className, _ := windows.UTF16PtrFromString("RandomizerClass")
	wc := WNDCLASSEXW{
		CbSize:        uint32(unsafe.Sizeof(WNDCLASSEXW{})),
		LpfnWndProc:   syscall.NewCallback(wndProc),
		HInstance:     windows.Handle(hInstance),
		HCursor:       windows.Handle(hCursor),
		LpszClassName: className,
	}

	ret, _, _ := procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))
	if ret == 0 {
		panic("Failed to register window class")
	}

	// Get screen dimensions for centering
	screenWidth, _, _ := procGetSystemMetrics.Call(SM_CXSCREEN)
	screenHeight, _, _ := procGetSystemMetrics.Call(SM_CYSCREEN)

	// Calculate center position
	posX := (int32(screenWidth) - windowSize) / 2
	posY := (int32(screenHeight) - windowSize) / 2

	// Create window (hidden from taskbar with WS_EX_TOOLWINDOW)
	windowName, _ := windows.UTF16PtrFromString("Randomizer")
	ret, _, _ = procCreateWindowExW.Call(
		WS_EX_TOPMOST|WS_EX_TOOLWINDOW,
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		WS_POPUP|WS_VISIBLE,
		uintptr(posX), uintptr(posY),
		uintptr(windowSize), uintptr(windowSize),
		0, 0,
		uintptr(hInstance),
		0,
	)
	hwnd = windows.Handle(ret)
	if hwnd == 0 {
		panic("Failed to create window")
	}

	// Show and update window explicitly
	procShowWindow.Call(uintptr(hwnd), SW_SHOW)
	procUpdateWindow.Call(uintptr(hwnd))
	procSetForegroundWindow.Call(uintptr(hwnd))

	// Set timer for number generation
	procSetTimer.Call(uintptr(hwnd), TIMER_ID, uintptr(currentIntervalSec*1000), 0)

	// Message loop
	var msg MSG
	for {
		ret, _, _ := procGetMessageW.Call(
			uintptr(unsafe.Pointer(&msg)),
			0, 0, 0,
		)
		if ret == 0 {
			break
		}
		procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		procDispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
	}
}

// Window procedure
func wndProc(hwnd windows.Handle, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case WM_DESTROY:
		procPostQuitMessage.Call(0)
		return 0

	case WM_SETCURSOR:
		// Always show arrow cursor (no hourglass)
		hCursor, _, _ := procLoadCursorW.Call(0, IDC_ARROW)
		procSetCursor.Call(hCursor)
		return 1

	case WM_ERASEBKGND:
		// Prevent flicker - we handle background in WM_PAINT
		return 1

	case WM_TIMER:
		// Generate new number on timer tick
		currentNumber = rand.Intn(100) + 1
		procInvalidateRect.Call(uintptr(hwnd), 0, 0)
		return 0

	case WM_PAINT:
		var ps PAINTSTRUCT
		hdc, _, _ := procBeginPaint.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&ps)))

		// Get client rect
		var rect RECT
		procGetClientRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect)))
		width := rect.Right - rect.Left
		height := rect.Bottom - rect.Top

		// Create memory DC for double buffering
		memDC, _, _ := procCreateCompatibleDC.Call(hdc)
		memBitmap, _, _ := procCreateCompatibleBitmap.Call(hdc, uintptr(width), uintptr(height))
		oldBitmap, _, _ := procSelectObject.Call(memDC, memBitmap)

		// Fill background with dark gray (RGB 22, 22, 22)
		brush, _, _ := procCreateSolidBrush.Call(0x00161616)
		procFillRect.Call(memDC, uintptr(unsafe.Pointer(&rect)), brush)
		procDeleteObject.Call(brush)

		// Calculate color based on current number (BGR format for Win32)
		// Apply inversion if needed
		displayNumber := currentNumber
		colorNumber := currentNumber
		if isInverted {
			colorNumber = 101 - currentNumber
		}
		r, g, b := calculateColor(colorNumber)
		textColor := uintptr(uint32(r) | (uint32(g) << 8) | (uint32(b) << 16))

		// Set text properties
		procSetTextColor.Call(memDC, textColor)
		procSetBkMode.Call(memDC, TRANSPARENT)

		// DYNAMIC AUTO-FIT: Calculate exact font size to fill 92% of window
		text := fmt.Sprintf("%d", displayNumber)
		textPtr, _ := windows.UTF16PtrFromString(text)
		fontName, _ := windows.UTF16PtrFromString("Arial Black")

		// Target size: 92% of window
		targetSize := float64(width) * 0.92

		// Step 1: Create test font with base height 100
		baseHeight := int32(100)
		testFont, _, _ := procCreateFontW.Call(
			uintptr(baseHeight),
			0, 0, 0,
			FW_BOLD,
			0, 0, 0,
			DEFAULT_CHARSET,
			OUT_TT_PRECIS,
			CLIP_DEFAULT_PRECIS,
			CLEARTYPE_QUALITY,
			VARIABLE_PITCH|FF_SWISS,
			uintptr(unsafe.Pointer(fontName)),
		)

		// Step 2: Measure text with test font
		oldTestFont, _, _ := procSelectObject.Call(memDC, testFont)
		var size SIZE
		procGetTextExtentPoint32W.Call(
			memDC,
			uintptr(unsafe.Pointer(textPtr)),
			uintptr(len(text)),
			uintptr(unsafe.Pointer(&size)),
		)

		// Step 3: Restore old font and delete test font
		procSelectObject.Call(memDC, oldTestFont)
		procDeleteObject.Call(testFont)

		// Step 4: Calculate scale factor (use minimum to fit both dimensions)
		scaleX := targetSize / float64(size.CX)
		scaleY := targetSize / float64(size.CY)
		scale := scaleX
		if scaleY < scaleX {
			scale = scaleY
		}

		// Step 5: Calculate final font height
		finalFontHeight := int32(float64(baseHeight) * scale)

		// Step 6: Create final font with calculated height
		hFont, _, _ := procCreateFontW.Call(
			uintptr(finalFontHeight),
			0, 0, 0,
			FW_BOLD,
			0, 0, 0,
			DEFAULT_CHARSET,
			OUT_TT_PRECIS,
			CLIP_DEFAULT_PRECIS,
			CLEARTYPE_QUALITY,
			VARIABLE_PITCH|FF_SWISS,
			uintptr(unsafe.Pointer(fontName)),
		)
		oldFont, _, _ := procSelectObject.Call(memDC, hFont)

		// Draw text centered (text and textPtr already declared above)
		procDrawTextW.Call(
			memDC,
			uintptr(unsafe.Pointer(textPtr)),
			uintptr(^uint(0)), // -1 as uintptr
			uintptr(unsafe.Pointer(&rect)),
			DT_CENTER|DT_VCENTER|DT_SINGLELINE|DT_NOPREFIX,
		)

		// Copy from memory DC to screen
		procBitBlt.Call(
			hdc,
			0, 0,
			uintptr(width), uintptr(height),
			memDC,
			0, 0,
			SRCCOPY,
		)

		// Cleanup GDI objects
		procSelectObject.Call(memDC, oldFont)
		procDeleteObject.Call(hFont)
		procSelectObject.Call(memDC, oldBitmap)
		procDeleteObject.Call(memBitmap)
		procDeleteDC.Call(memDC)

		procEndPaint.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&ps)))
		return 0

	case WM_LBUTTONDOWN:
		// Use standard Windows drag technique
		procReleaseCapture.Call()
		procSendMessageW.Call(uintptr(hwnd), WM_NCLBUTTONDOWN, HTCAPTION, 0)
		return 0

	case WM_RBUTTONDOWN:
		// Start tracking for click vs drag
		isRDown = true
		isDragging = false
		var pt POINT
		procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
		resizeStartX = pt.X
		resizeStartY = pt.Y
		resizeStartSize = windowSize
		procSetCapture.Call(uintptr(hwnd))
		return 0

	case WM_RBUTTONUP:
		isRDown = false
		procReleaseCapture.Call()
		// If not dragging, show context menu
		if !isDragging {
			showContextMenu(hwnd)
		}
		isDragging = false
		return 0

	case WM_MOUSEMOVE:
		if isRDown {
			// Check if mouse moved enough to be considered dragging
			var pt POINT
			procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
			dx := pt.X - resizeStartX
			dy := pt.Y - resizeStartY

			// Drag threshold: 4 pixels
			if !isDragging && (abs(dx) > 4 || abs(dy) > 4) {
				isDragging = true
			}

			if isDragging {
				// Resize window smoothly
				newSize := resizeStartSize + dx

				// Clamp size between 35 and 500
				if newSize < 35 {
					newSize = 35
				} else if newSize > 500 {
					newSize = 500
				}

				windowSize = newSize

				// Get current position
				var rect RECT
				procGetWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect)))

				// Resize window (keeping it square)
				procSetWindowPos.Call(
					uintptr(hwnd),
					0,
					uintptr(rect.Left), uintptr(rect.Top),
					uintptr(windowSize), uintptr(windowSize),
					SWP_NOZORDER,
				)

				// Force immediate redraw
				procInvalidateRect.Call(uintptr(hwnd), 0, 0)
				procUpdateWindow.Call(uintptr(hwnd))
			}
		}
		return 0

	case WM_MBUTTONUP:
		// Middle mouse button - exit
		procPostQuitMessage.Call(0)
		return 0

	case WM_KEYDOWN:
		if wParam == VK_ESCAPE {
			procPostQuitMessage.Call(0)
			return 0
		}

	case WM_COMMAND:
		cmdID := uint32(wParam & 0xFFFF)
		handleMenuCommand(hwnd, cmdID)
		return 0
	}

	ret, _, _ := procDefWindowProcW.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
	return ret
}

// showContextMenu creates and displays context menu
func showContextMenu(hwnd windows.Handle) {
	// Create main popup menu
	hMenu, _, _ := procCreatePopupMenu.Call()
	if hMenu == 0 {
		return
	}

	// Create Interval submenu
	hIntervalMenu, _, _ := procCreatePopupMenu.Call()
	addMenuItem(hIntervalMenu, CMD_INTERVAL_1S, "1s", currentIntervalSec == 1)
	addMenuItem(hIntervalMenu, CMD_INTERVAL_3S, "3s", currentIntervalSec == 3)
	addMenuItem(hIntervalMenu, CMD_INTERVAL_5S, "5s", currentIntervalSec == 5)
	addMenuItem(hIntervalMenu, CMD_INTERVAL_7S, "7s", currentIntervalSec == 7)
	addMenuItem(hIntervalMenu, CMD_INTERVAL_10S, "10s", currentIntervalSec == 10)

	// Create Color submenu
	hColorMenu, _, _ := procCreatePopupMenu.Call()
	addMenuItem(hColorMenu, CMD_PALETTE_NORM, "Green → Red", !isInverted)
	addMenuItem(hColorMenu, CMD_PALETTE_INV, "Red → Green", isInverted)

	// Add submenus to main menu
	intervalText, _ := windows.UTF16PtrFromString("Interval")
	procAppendMenuW.Call(hMenu, MF_POPUP, hIntervalMenu, uintptr(unsafe.Pointer(intervalText)))

	colorText, _ := windows.UTF16PtrFromString("Color")
	procAppendMenuW.Call(hMenu, MF_POPUP, hColorMenu, uintptr(unsafe.Pointer(colorText)))

	// Add separator
	procAppendMenuW.Call(hMenu, MF_SEPARATOR, 0, 0)

	// Add Exit
	exitText, _ := windows.UTF16PtrFromString("Exit")
	procAppendMenuW.Call(hMenu, MF_STRING, CMD_EXIT, uintptr(unsafe.Pointer(exitText)))

	// Get cursor position
	var pt POINT
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))

	// Show menu
	procTrackPopupMenuEx.Call(
		hMenu,
		TPM_RIGHTBUTTON|TPM_TOPALIGN,
		uintptr(pt.X), uintptr(pt.Y),
		uintptr(hwnd),
		0,
	)

	// Cleanup menu handles
	procDestroyMenu.Call(hMenu)
}

// addMenuItem adds a menu item with optional check mark
func addMenuItem(hMenu uintptr, cmdID uint32, text string, checked bool) {
	textPtr, _ := windows.UTF16PtrFromString(text)
	flags := uint32(MF_STRING)
	if checked {
		flags |= MF_CHECKED
	}
	procAppendMenuW.Call(hMenu, uintptr(flags), uintptr(cmdID), uintptr(unsafe.Pointer(textPtr)))
}

// handleMenuCommand processes menu commands
func handleMenuCommand(hwnd windows.Handle, cmdID uint32) {
	switch cmdID {
	case CMD_INTERVAL_1S:
		setInterval(hwnd, 1)
	case CMD_INTERVAL_3S:
		setInterval(hwnd, 3)
	case CMD_INTERVAL_5S:
		setInterval(hwnd, 5)
	case CMD_INTERVAL_7S:
		setInterval(hwnd, 7)
	case CMD_INTERVAL_10S:
		setInterval(hwnd, 10)
	case CMD_PALETTE_NORM:
		isInverted = false
		procInvalidateRect.Call(uintptr(hwnd), 0, 1)
	case CMD_PALETTE_INV:
		isInverted = true
		procInvalidateRect.Call(uintptr(hwnd), 0, 1)
	case CMD_EXIT:
		procDestroyWindow.Call(uintptr(hwnd))
	}
}

// setInterval updates the timer interval
func setInterval(hwnd windows.Handle, seconds uint32) {
	currentIntervalSec = seconds
	procKillTimer.Call(uintptr(hwnd), TIMER_ID)
	procSetTimer.Call(uintptr(hwnd), TIMER_ID, uintptr(seconds*1000), 0)
}

// abs returns absolute value of int32
func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// calculateColor computes RGB color based on number (1-100)
// Green (1) -> Lime (25) -> Yellow (50) -> Orange (75) -> Red (100)
func calculateColor(n int) (r, g, b uint8) {
	t := float64(n-1) / 99.0

	if t <= 0.5 {
		// Green to Yellow (0,255,0) -> (255,255,0)
		r = uint8(255 * t * 2.0)
		g = 255
		b = 0
	} else {
		// Yellow to Red (255,255,0) -> (255,0,0)
		r = 255
		g = uint8(255 * (2.0 - t*2.0))
		b = 0
	}

	return r, g, b
}
