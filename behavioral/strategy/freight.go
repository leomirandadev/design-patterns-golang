package main

type freightStrategy interface {
	getPrice() float32
}

func newSedex(distanceKm float32) freightStrategy {
	return &sedex{distanceKm: distanceKm}
}

type sedex struct {
	distanceKm float32
}

func (s *sedex) getPrice() float32 {
	var basePrice float32 = 10

	if s.distanceKm > 500 {
		basePrice = 30
	}

	return basePrice
}
