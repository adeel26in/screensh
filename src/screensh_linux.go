func linux() {

	fmt.Println("Welcome to screensh!")

	fmt.Println("REQUIRES GNOME-SCREENSHOT!")

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

	cmd := exec.Command("gnome-screenshot", "--include-pointer", "--delay", "2", "-f", fullpath)

	err := cmd.Run()

	if err != nil {

		fmt.Println(err)
		return
	}
}
