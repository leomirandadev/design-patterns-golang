package main

import (
	"fmt"
)

func main() {
	err := createAccountFacade(&user{username: "Leo", role: 1, password: "12312312"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("sucesso")

}

// two responsabilities but simplify client call
func createAccountFacade(user *user) error {

	if err := createUser(user); err != nil {
		return err
	}

	if err := createAccount(user.id); err != nil {
		return err
	}

	return nil
}
