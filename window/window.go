package window

import (
	"syscall"
	"unsafe"
)

func FindWindows(mod *syscall.LazyDLL) uintptr {
	proc := mod.NewProc("FindWindowW")

	className := syscall.StringToUTF16Ptr("Notepad")

	hWnd, _, err := proc.Call(uintptr(unsafe.Pointer(className)), 0)
	if hWnd == 0 {
		panic(err)
	}

	return hWnd
}

func SetWindowPosition(mod *syscall.LazyDLL, hWnd uintptr, x, y, width, height int32) {
	proc := mod.NewProc("SetWindowPos")

	proc.Call(
		hWnd,
		uintptr(0), // Z-Order
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		0x0040, // 플래그 SWP_SHOWWINDOW
	)
}

func ShowWindow(mod *syscall.LazyDLL)
