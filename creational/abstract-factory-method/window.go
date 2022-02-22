package main

import "fmt"

type windowsButton struct{}
type windowsTextField struct{}

func createWindowsButton() buttonInterface {
	return windowsButton{}
}

func (windowsButton) click() {
	fmt.Println("click windows!")
}

func createWindowsTextField() textFieldInterface {
	return windowsTextField{}
}

func (windowsTextField) typing(value string) {
	fmt.Println("typing:", value, "text field windows!")
}
