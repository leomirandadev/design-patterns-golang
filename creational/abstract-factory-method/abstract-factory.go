package main

type osFactory interface {
	createButton() buttonInterface
	createTextField() textFieldInterface
}

type os struct {
	typeOs string
}

func createOs(typeOs string) osFactory {
	return os{typeOs: typeOs}
}

func (o os) createButton() buttonInterface {
	switch o.typeOs {
	case "linux":
		return createLinuxButton()
	case "windows":
		return createWindowsButton()
	default:
		return nil
	}
}

func (o os) createTextField() textFieldInterface {
	switch o.typeOs {
	case "linux":
		return createLinuxTextField()
	case "windows":
		return createWindowsTextField()
	default:
		return nil
	}
}
