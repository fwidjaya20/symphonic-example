package cloud_messaging

import (
	"github.com/fwidjaya20/symphonic-example/src/notification/domain/entity"
)

type CloudMessaging interface {
	Send(msg entity.Message) error
}
