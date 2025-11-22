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
	case "freebsd":
		freebsd()
	default:
		fmt.Println("Unsupported OS: ", useros)
	}
}
