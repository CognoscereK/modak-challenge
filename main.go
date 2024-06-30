package main

import (
	"time"
)

func main() {
	rateLimiter := NewInMemoryRateLimiter()
	gateway := NewPrintGateway()

	rateLimiter.SetLimit("Status", RateLimit{Count: 2, Interval: time.Minute})
	rateLimiter.SetLimit("News", RateLimit{Count: 1, Interval: 24 * time.Hour})
	rateLimiter.SetLimit("Marketing", RateLimit{Count: 3, Interval: time.Hour})

	notificationService := NewNotificationService(rateLimiter, gateway)

	notifications := []Notification{
		{"Status", "foobar@gmail.com", "Status Update 1"},
		{"Status", "foobar@gmail.com", "Status Update 2"},
		{"Status", "foobar@gmail.com", "Status Update 3"}, // rate exceed
		{"News", "foobar@gmail.com", "Daily News 1"},
		{"News", "foobar@gmail.com", "Daily News 2"}, // rate exceeed
		{"News", "foo@gmail.com", "Daily News 1"},
		{"Marketing", "foobar@gmail.com", "Marketing message 1"},
		{"Marketing", "foobar@gmail.com", "Marketing message 2"},
		{"Marketing", "foobar@gmail.com", "Marketing message 3"},
		{"Marketing", "foobar@gmail.com", "Marketing message 4"}, // rate exceed
	}

	for _, notification := range notifications {
		notificationService.SendNotification(notification)
	}
}
