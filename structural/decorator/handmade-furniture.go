package main

type IHandmadeFurniture interface {
	getPrice() float32
	getName() string
	whoMade() string
}

func newHandmadeFurniture(f IFurniture) IHandmadeFurniture {
	return handmadeFurniture{furniture: f}
}

type handmadeFurniture struct {
	furniture IFurniture
}

func (hf handmadeFurniture) getPrice() float32 {
	return hf.furniture.getPrice() + 10
}

func (hf handmadeFurniture) getName() string {
	return hf.furniture.getName() + " made for " + hf.whoMade()
}

func (handmadeFurniture) whoMade() string {
	return "Thomas"
}
