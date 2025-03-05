package main

import (
	"fmt"

	"syscall"
	"win_api/cursor"
)

func main() {
	mod := syscall.NewLazyDLL("user32.dll")

	pt, err := cursor.GetCursorPosition(mod)
	if err != nil {
		panic(err)
	}

	fmt.Println(pt)
}
