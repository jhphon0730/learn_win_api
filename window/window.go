package window

import (
	"fmt"
	"syscall"
	"unsafe"
)

func FindWindows(mod *syscall.LazyDLL) {
	proc := mod.NewProc("FindWindowW")

	className := syscall.StringToUTF16Ptr("Notepad")

	hWnd, _, _ := proc.Call(uintptr(unsafe.Pointer(className)), 0)

	fmt.Println(hWnd)
}
