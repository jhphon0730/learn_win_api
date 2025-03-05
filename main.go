package main

import (
	"fmt"

	"syscall"
	"win_api/cursor"
	"win_api/message"
	"win_api/metrics"
	"win_api/window"
)

func callMousePointer() {
	mod := syscall.NewLazyDLL("user32.dll")

	pt, err := cursor.GetCursorPosition(mod)
	if err != nil {
		panic(err)
	}

	fmt.Println(pt)
}

func callMessageBox() {
	mod := syscall.NewLazyDLL("user32.dll")

	message.GetMessageBox(mod)
}

func callSystemMetrics() {
	mod := syscall.NewLazyDLL("user32.dll")

	metrics.GetSystemMetrics(mod)
}

func callFindWindows() {
	mod := syscall.NewLazyDLL("user32.dll")

	window.FindWindows(mod)
}

func callSetWindowPosition() {
	mod := syscall.NewLazyDLL("user32.dll")

	hWnd := window.FindWindows(mod)

	window.SetWindowPosition(mod, hWnd, 0, 0, 800, 600)
}

func callShowWindow() {
	mod := syscall.NewLazyDLL("user32.dll")

	hWnd := window.FindWindows(mod)

	window.ShowWindow(mod, hWnd, 3)
}

func main() {
	callShowWindow()
}
