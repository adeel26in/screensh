package main

import (
	"fmt"
	"runtime"
)

func main() {

	useros := runtime.GOOS

	switch useros {
	case "linux":
		linux()
	case "darwin":
		darwin()
	case "windows":
		windows()
	default:
		fmt.Println("Couldn't detect your OS!")
	}
}
