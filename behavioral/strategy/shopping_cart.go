package main

type IShoppingCart interface {
	add(product product)
	getTotal(freight freightStrategy) float32
}

func newShoppingCart() IShoppingCart {
	return &shoppingCart{}
}

type shoppingCart struct {
	products []product
}

func (s *shoppingCart) add(product product) {
	s.products = append(s.products, product)
}

func (s *shoppingCart) getTotal(freight freightStrategy) float32 {
	var total float32
	for _, product := range s.products {
		total += product.value
	}

	total += freight.getPrice()

	return total
}
