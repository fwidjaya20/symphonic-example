package application

import "github.com/fwidjaya20/symphonic-skeleton/src/notification/application/command"

type NotificationApplication struct {
	Send command.SendHandler
}

func NewNotificationApplication() NotificationApplication {
	return NotificationApplication{
		Send: command.NewSendHandler(),
	}
}
