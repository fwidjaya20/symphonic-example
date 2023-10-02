package cloud_messaging

import (
	NotificationEnum "github.com/fwidjaya20/symphonic-skeleton/shared/enum/notification"
	"github.com/fwidjaya20/symphonic-skeleton/src/notification/domain/entity"
)

type CloudMessaging interface {
	Send(msg entity.Message) error
}

func GetCloudMessaging(channel NotificationEnum.Channel) CloudMessaging {
	switch channel {
	case NotificationEnum.FirebaseCloudMessaging:
		return NewFirebaseCloudMessaging()
	default:
		return nil
	}
}
