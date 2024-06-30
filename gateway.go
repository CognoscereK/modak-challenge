package main

import "log"

type Gateway interface {
	Send(notification Notification)
}

type PrintGateway struct{}

func NewPrintGateway() *PrintGateway {
	return &PrintGateway{}
}

func (gateway *PrintGateway) Send(notification Notification) {
	log.Printf("Sending notification of type %s to %s with content: %s", notification.Type, notification.Recipient, notification.Message)
}
