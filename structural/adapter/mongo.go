package main

import "fmt"

func newMongo() mongoImplement {
	return mongoImplement{}
}

type mongoImplement struct{}

func (mongoImplement) insertOne() {
	fmt.Println("insert mongo")
}
