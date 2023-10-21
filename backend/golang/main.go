package main

import (
	"fmt"
)

type Notifier interface {
	Send(msg string)
}
type notificationService struct {
	notifier Notifier
}

type EmailNotifier struct {
	notifierType string
}

func (n EmailNotifier) Send(msg string) {
	// Do send email logic
	fmt.Printf("Sending message: %s (Sender: %s)\n", msg, n.notifierType)
}

type SmsNotifier struct {
	notifierType string
}

func (n SmsNotifier) Send(msg string) {
	// Do send sms logic
	fmt.Printf("Sending message: %s (Sender: %s)\n", msg, n.notifierType)
}

func main() {
	smsNotifier := SmsNotifier{"sms"}
	emailNotifier := EmailNotifier{"email"}
	s := notificationService{smsNotifier}
	s.notifier.Send("Hello World")

	s = notificationService{emailNotifier}
	s.notifier.Send("Hello World")
}
