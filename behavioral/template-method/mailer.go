package main

type mailerProvider interface {
	validEmailSender(email, password string) error
	sendEmail() error
	getNameProvider() string
}
