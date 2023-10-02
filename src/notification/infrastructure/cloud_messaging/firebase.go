package cloud_messaging

import (
	"fmt"

	"github.com/fwidjaya20/symphonic-skeleton/src/notification/domain/entity"
)

type FCM struct{}

func NewFirebaseCloudMessaging() CloudMessaging {
	return &FCM{}
}

func (f *FCM) Send(msg entity.Message) error {
	fmt.Printf("[FCM][%s][To: %s] %s\n", msg.Subject, msg.Recipient, msg.Body)
	return nil
}
