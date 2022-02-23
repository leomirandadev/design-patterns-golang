package main

import (
	"errors"
	"fmt"
)

func newHotmail() mailerProvider {
	return hotmail{}
}

type hotmail struct{}

func (hotmail) getNameProvider() string {
	return "hotmail"
}

func (hotmail) validEmailSender(email, password string) error {
	fmt.Println("hotmail implementation validEmailSender")
	return errors.New("email not valid")
}

func (hotmail) sendEmail() error {
	fmt.Println("hotmail implementation sendEmail")
	return errors.New("email not send")
}
