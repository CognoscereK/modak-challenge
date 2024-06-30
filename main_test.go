package main

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	rateLimiter := NewInMemoryRateLimiter()

	rateLimiter.SetLimit("Status", RateLimit{Count: 2, Interval: time.Minute})
	rateLimiter.SetLimit("News", RateLimit{Count: 1, Interval: 24 * time.Hour})

	tests := []struct {
		notification Notification
		expectSend   bool
	}{
		{Notification{"Status", "user1@example.com", "Status update 1"}, true},
		{Notification{"Status", "user1@example.com", "Status update 2"}, true},
		{Notification{"Status", "user1@example.com", "Status update 3"}, false},
		{Notification{"News", "user1@example.com", "Daily news 1"}, true},
		{Notification{"News", "user1@example.com", "Daily news 2"}, false},
	}

	for _, test := range tests {
		canSend := rateLimiter.CanSend(test.notification)
		if canSend != test.expectSend {
			t.Errorf("expected %v, got %v for notification: %+v", test.expectSend, canSend, test.notification)
		}
	}
}

type SpyGateway struct {
	CallsPerRecipient map[string]int
}

func (gateway *SpyGateway) Send(notification Notification) {
	gateway.CallsPerRecipient[notification.Recipient]++
}

func TestNotificationService(t *testing.T) {
	rateLimiter := NewInMemoryRateLimiter()
	rateLimiter.SetLimit("Status", RateLimit{Count: 2, Interval: time.Minute})
	rateLimiter.SetLimit("News", RateLimit{Count: 1, Interval: 24 * time.Hour})

	spyGateway := &SpyGateway{CallsPerRecipient: make(map[string]int)}

	notificationService := NewNotificationService(rateLimiter, spyGateway)

	tests := []struct {
		notification Notification
	}{
		{Notification{"Status", "user1@example.com", "Status update 1"}},
		{Notification{"Status", "user1@example.com", "Status update 2"}},
		{Notification{"Status", "user1@example.com", "Status update 3"}},
		{Notification{"News", "user1@example.com", "Daily news 1"}},
		{Notification{"News", "user1@example.com", "Daily news 2"}},
		{Notification{"News", "user2@example.com", "Daily news 1"}},
	}

	for _, test := range tests {
		notificationService.SendNotification(test.notification)
	}

	if spyGateway.CallsPerRecipient["user1@example.com"] != 3 {
		t.Error("Expected gateway to be called 3 times for user1@example.com")
	}

	if spyGateway.CallsPerRecipient["user2@example.com"] != 1 {
		t.Error("Expected gateway to be called 1 time for user2@example.com")
	}
}
