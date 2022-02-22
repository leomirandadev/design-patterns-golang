package main

import "fmt"

func createUser(user *user) error {
	fmt.Println("creating user on db")
	user.id = "123"
	return nil
}

func createAccount(userID string) error {
	fmt.Println("creating account on db")
	return nil
}
