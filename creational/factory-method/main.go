package main

import (
	"errors"
	"fmt"
)

func main() {
	transport, err := factoryTransport("car")
	if err == nil {
		transport.goToDestination()
	}

	transport2, err := factoryTransport("truck")
	if err == nil {
		transport2.goToDestination()
	}

	transport3, err := factoryTransport("")
	if err == nil {
		transport3.goToDestination()
	}
}

func factoryTransport(transportType string) (transportInterface, error) {
	switch transportType {
	case "car":
		return newCar(), nil
	case "truck":
		return newTruck(), nil
	default:
		return nil, errors.New("we can't create transport")
	}
}

type transportInterface interface {
	goToDestination()
}

func newTruck() transportInterface {
	return truck{}
}

type truck struct{}

func (truck) goToDestination() {
	fmt.Println("truck going")
}

func newCar() transportInterface {
	return car{}
}

type car struct{}

func (car) goToDestination() {
	fmt.Println("car going")
}
