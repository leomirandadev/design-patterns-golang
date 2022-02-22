package main

import "fmt"

type linuxButton struct{}
type linuxTextField struct{}

func createLinuxButton() buttonInterface {
	return linuxButton{}
}

func (linuxButton) click() {
	fmt.Println("click linux!")
}

func createLinuxTextField() textFieldInterface {
	return linuxTextField{}
}

func (linuxTextField) typing(value string) {
	fmt.Println("typing:", value, "text field linux!")
}
