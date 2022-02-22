package main

type IFurniture interface {
	getPrice() float32
	getName() string
}

func newFurniture(name string) IFurniture {
	return furniture{name: name}
}

type furniture struct {
	name string
}

func (furniture) getPrice() float32 {
	return 30.0
}

func (f furniture) getName() string {
	return f.name
}
