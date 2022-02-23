package main

import "fmt"

type IObserver interface {
	update(titleNews string)
	getID() string
}

func newUser(id, name string) IObserver {
	return user{name: name, id: id}
}

type user struct {
	id   string
	name string
}

func (u user) update(titleNews string) {
	fmt.Println("Send message to ", u.name, " telling the name of news is ", titleNews)
}

func (u user) getID() string {
	return u.id
}
