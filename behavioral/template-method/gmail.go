package main

import "fmt"

func newGmail() mailerProvider {
	return gmail{}
}

type gmail struct{}

func (gmail) getNameProvider() string {
	return "gmail"
}

func (gmail) validEmailSender(email, password string) error {
	fmt.Println("gmail implementation validEmailSender")
	return nil
}

func (gmail) sendEmail() error {
	fmt.Println("gmail implementation sendEmail")
	return nil
}
