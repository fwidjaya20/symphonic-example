package factory

import (
	NotificationEnum "github.com/fwidjaya20/symphonic-example/shared/enum/notification"
	DomainInterface "github.com/fwidjaya20/symphonic-example/src/notification/domain/interface/cloud_messaging"
	"github.com/fwidjaya20/symphonic-example/src/notification/infrastructure/cloud_messaging"
)

func GetCloudMessaging(channel NotificationEnum.Channel) DomainInterface.CloudMessaging {
	switch channel {
	case NotificationEnum.FirebaseCloudMessaging:
		return cloud_messaging.NewFirebaseCloudMessaging()
	default:
		return nil
	}
}
