package main

import "fmt"

func main() {

	f := newFurniture("table")

	fmt.Println(f.getName())
	fmt.Println("R$", f.getPrice())

	hf := newHandmadeFurniture(f)

	fmt.Println(hf.getName())
	fmt.Println("R$", hf.getPrice())
}
