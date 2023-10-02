package public

import NotificationEnum "github.com/fwidjaya20/symphonic-example/shared/enum/notification"

type SendNotificationRequest struct {
	Channel NotificationEnum.Channel `json:"channel"`
	From    string                   `json:"from"`
	To      string                   `json:"to"`
	Subject string                   `json:"subject"`
	Body    string                   `json:"body"`
}
