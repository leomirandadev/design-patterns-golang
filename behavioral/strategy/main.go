package main

import "fmt"

func main() {

	sc := newShoppingCart()
	sc.add(product{name: "shirt", value: 30})
	sc.add(product{name: "pants", value: 100})

	var distanceKm float32 = 1000

	fr := newSedex(distanceKm)

	fmt.Println(sc.getTotal(fr))
}
