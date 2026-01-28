package main

import "fmt"

func main() {
	email := EmailNotifier{"test@mail.ru"}
	err := email.Send("email mess")
	if err != nil {
		fmt.Println(err)
	}
	sms := SMSNotifier{"444877"}
	err = sms.Send("sms mess")
	if err != nil {
		fmt.Println(err)
	}
}

type Notifier interface { //КОД ДНЯ 1154
	Send(message string) error
}

type EmailNotifier struct {
	Email string
}

func (n EmailNotifier) Send(message string) error {
	fmt.Printf("Sending email to %s\n", n.Email)
	return nil
}

type SMSNotifier struct {
	PhoneNumber string
}

func (n SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS to %s\n", n.PhoneNumber)
	return nil
}
