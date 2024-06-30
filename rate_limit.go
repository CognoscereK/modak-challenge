package main

import (
	"time"

	"github.com/samber/lo"
)

type RateLimit struct {
	Count    int
	Interval time.Duration
}

type RateLimiter interface {
	SetLimit(notificationType NotificationType, limit RateLimit)
	CanSend(notification Notification) bool
}

type InMemoryRateLimiter struct {
	limits            map[NotificationType]RateLimit
	sentNotifications map[string]map[NotificationType][]time.Time
}

func NewInMemoryRateLimiter() *InMemoryRateLimiter {
	return &InMemoryRateLimiter{
		limits:            make(map[NotificationType]RateLimit),
		sentNotifications: make(map[string]map[NotificationType][]time.Time),
	}
}

func (ratelimit *InMemoryRateLimiter) SetLimit(notificationType NotificationType, limit RateLimit) {
	ratelimit.limits[notificationType] = limit
}

func (ratelimit *InMemoryRateLimiter) CanSend(notification Notification) bool {
	// Initialize sent notifications submap to avoid errors
	if ratelimit.sentNotifications[notification.Recipient] == nil {
		ratelimit.sentNotifications[notification.Recipient] = make(map[NotificationType][]time.Time)
	}

	limit, exists := ratelimit.limits[notification.Type]

	// We always can send if there is no rate limit defined
	if !exists {
		return true
	}

	now := time.Now().UTC()

	// we count how many nofications were sent in the defined limit interval
	count := lo.CountBy(ratelimit.sentNotifications[notification.Recipient][notification.Type], func(t time.Time) bool {
		return now.Sub(t) < limit.Interval
	})

	// if the number of notifications sent plus the one that we are sending now do not exceed the rate limit it means we can send it otherwise we just return false
	if count+1 <= limit.Count {
		ratelimit.sentNotifications[notification.Recipient][notification.Type] = append(ratelimit.sentNotifications[notification.Recipient][notification.Type], time.Now().UTC())

		return true
	}

	return false
}
