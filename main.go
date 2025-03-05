package main

import (
	"fmt"

	"syscall"
	"win_api/cursor"
	"win_api/message"
	"win_api/metrics"
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

func main() {

}
