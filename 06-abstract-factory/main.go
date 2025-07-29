package main

import "fmt"

// 1. Interfaces
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// 2. Structs
// 2.1 SMS Notification
type SMSNotification struct{}

func (SMSNotification) SendNotification() {
	fmt.Print("📱 Sending SMS notification")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// 2.2 Email Notification
type EmailNotification struct{}

func (EmailNotification) SendNotification() {
	fmt.Print("📧 Sending Email notification")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

// 3. Sender
// 3.1 SMS Sender
type SMSNotificationSender struct{}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Mobile"
}

// 3.2 Email Sender
type EmailNotificationSender struct{}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "Internet"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &SMSNotification{}, nil
	case "Email":
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("Unknown notification type: %s", notificationType)
}

func sendNotification(factory INotificationFactory) {
	factory.SendNotification()
}

func getMethod(factory INotificationFactory) {
	fmt.Println("\nSender Method:", factory.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
