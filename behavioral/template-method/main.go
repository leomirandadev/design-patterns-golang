package main

import "fmt"

func main() {
	mailerProvider := newGmail()

	mailer := newMailer(mailerProvider)
	if err := mailer.send("leonardo@gmail.com", "123"); err != nil {
		fmt.Println(err)
	}

}
