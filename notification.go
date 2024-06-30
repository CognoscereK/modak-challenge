package main

type NotificationType string

type Notification struct {
	Type      NotificationType
	Recipient string
	Message   string
}
