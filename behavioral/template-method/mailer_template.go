package main

import (
	"fmt"
)

func newMailer(mailerProvider mailerProvider) mailerTemplateMethod {
	return mailerTemplateMethod{mailerProvider: mailerProvider}
}

type mailerTemplate interface {
	send(email, password string)
}

type mailerTemplateMethod struct {
	mailerProvider mailerProvider
}

func (e mailerTemplateMethod) send(email, password string) error {

	if err := e.mailerProvider.validEmailSender(email, password); err != nil {
		return err
	}

	if err := e.mailerProvider.sendEmail(); err != nil {
		return err
	}

	fmt.Println("email sended by", e.mailerProvider.getNameProvider())
	return nil
}
