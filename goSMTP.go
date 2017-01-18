package main

import (

	"log"
	"net/smtp"
)

func main() {
	send("Apelle", "receiver")
}

func send(body string, to string) {
	from := "email"
	password := "password"

	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: Your messages subject" + "\r\n\r\n" +
		body + "\r\n"

	err := smtp.SendMail("smtpserver:port", smtp.PlainAuth("", from, password, "smtpserver"), from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	log.Print("message sent")
}