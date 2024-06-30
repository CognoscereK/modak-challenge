package main

import "log"

type NotificationService struct {
	rateLimiter RateLimiter
	gateway     Gateway
}

func NewNotificationService(rateLimiter RateLimiter, gateway Gateway) *NotificationService {
	return &NotificationService{
		rateLimiter: rateLimiter,
		gateway:     gateway,
	}
}

func (notificationService *NotificationService) SendNotification(notification Notification) {
	// we only send the notification if the rate limit allows it

	if notificationService.rateLimiter.CanSend(notification) {
		notificationService.gateway.Send(notification)
	} else {
		log.Printf("Rate limit exceeed for %s and type %s", notification.Recipient, notification.Type)
	}
}
