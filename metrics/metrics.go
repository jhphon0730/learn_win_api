package metrics

import (
	"syscall"
)

func GetSystemMetrics(mod *syscall.LazyDLL) {
	proc := mod.NewProc("GetSystemMetrics")

	width, _, _ := proc.Call(uintptr(0))
	height, _, _ := proc.Call(uintptr(1))

	println("Width: ", width)
	println("Height: ", height)
}
