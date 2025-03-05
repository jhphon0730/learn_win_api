package cursor

import (
	"syscall"
	"unsafe"
)

type Point struct {
	X, Y int32
}

func GetCursorPosition(mod *syscall.LazyDLL) (*Point, error) {
	proc := mod.NewProc("GetCursorPos")

	point := &Point{}

	ret, _, err := proc.Call(uintptr(unsafe.Pointer(point)))
	if ret == 0 {
		return nil, err
	}

	return point, nil
}
