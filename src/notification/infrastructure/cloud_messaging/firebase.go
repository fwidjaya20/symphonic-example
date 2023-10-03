package cloud_messaging

import (
	"fmt"

	"github.com/fwidjaya20/symphonic-example/src/notification/domain/entity"
	DomainInterface "github.com/fwidjaya20/symphonic-example/src/notification/domain/interface/cloud_messaging"
)

type FCM struct{}

func NewFirebaseCloudMessaging() DomainInterface.CloudMessaging {
	return &FCM{}
}

func (f *FCM) Send(msg entity.Message) error {
	fmt.Printf("[FCM][%s][To: %s] %s\n", msg.Subject, msg.Recipient, msg.Body)
	return nil
}
