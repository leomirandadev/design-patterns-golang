package main

// abstract factory is
func main() {
	windows := createOs("windows")
	windows.createTextField().typing("leonardo")
	windows.createButton().click()

	linux := createOs("linux")
	linux.createTextField().typing("nicole")
	linux.createButton().click()
}
