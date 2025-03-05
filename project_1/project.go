package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	ptr *POINT = &POINT{}
)

type POINT struct {
	X, Y int32
}

func main() {
	// step 0: message box
	messageProc := user32.NewProc("MessageBoxW")

	// step 1: get the currnt cursor position
	proc := user32.NewProc("GetCursorPos")
	ret, _, err := proc.Call(uintptr(unsafe.Pointer(ptr)))
	if ret == 0 {
		panic(err)
	}

	// step 2: get the current window handle
	proc = user32.NewProc("GetSystemMetrics")
	system_width, _, _ := proc.Call(uintptr(0))
	system_height, _, _ := proc.Call(uintptr(1))

	// step 3: get the current window handle
	proc = kernel32.NewProc("GetTickCount")
	tickCount, _, _ := proc.Call()
	uptimeMessage := float64(tickCount) / 1000

	messageProc.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fmt.Sprintf("x=%d, y=%d", ptr.X, ptr.Y)))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Mouse Position"))),
		uintptr(0),
	)

	messageProc.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fmt.Sprintf("width=%d, height=%d", system_width, system_height)))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("system metrics"))),
		uintptr(0),
	)

	messageProc.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fmt.Sprintf("uptime=%v", uptimeMessage)))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("system uptime"))),
		uintptr(0),
	)
}
