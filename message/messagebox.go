package message

import (
	"syscall"
	"unsafe"
)

func GetMessageBox(mod *syscall.LazyDLL) {
	proc := mod.NewProc("MessageBoxW")

	proc.Call(
		uintptr(0), // 부모 윈도우 핸들 ( null == 0)
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Hello World, message here"))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("window title here"))),
		uintptr(0), // 버튼 종류 ( 0 == OK )
	)
}
