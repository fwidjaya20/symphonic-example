package event

import (
	SharedContext "github.com/fwidjaya20/symphonic-example/shared/context"
	CalendarListener "github.com/fwidjaya20/symphonic-example/src/calendar/application/listener"
	NotificationListener "github.com/fwidjaya20/symphonic-example/src/notification/application/listener"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/event/task"
	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/labstack/echo/v4"
)

type Kernel struct {
}

func (k *Kernel) Listeners() ContractEvent.Collection {
	e := echo.New()

	ctx := &SharedContext.SymphonicContext{
		Context: e.NewContext(nil, nil),
	}

	return ContractEvent.Collection{
		task.TaskCreated{}.Signature(): []ContractEvent.Listener{
			CalendarListener.NewTaskCreatedSubscriber(ctx),
			NotificationListener.NewTaskCreatedSubscriber(ctx),
		},
	}
}
