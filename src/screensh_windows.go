//go:build windows

package main

import (
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/adeel26in/screenshot"
)

func windows() {
	fmt.Println("Welcome to screensh!")

	fmt.Println("SCREENSHOTS AFTER 2 SECONDS!")

	var fileaddress string
	var filename string
	fmt.Println("EXAMPLE: /home/<user>/examplefolder/")

	fmt.Print("Where do you want to save the file?: ")
	fmt.Scanln(&fileaddress)

	fmt.Println("EXAMPLE (ONLY PNG): example.png")

	fmt.Print("What do you want to name the file?: ")
	fmt.Scanln(&filename)

	fullpath := filepath.Join(fileaddress, filename)

	time.Sleep(2 * time.Second)
	
	actualimage, err := screenshot.CaptureScreen()

	if err != nil {

		fmt.Println(err)
		return
	}

	filecreate, err := os.Create(fullpath)
	if err != nil {

		fmt.Println(err)
		return
	}

	defer filecreate.Close()

	err = png.Encode(filecreate, actualimage)

	if err != nil {

		fmt.Println(err)
		return
	}
}
