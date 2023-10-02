package event

import (
	SharedContext "github.com/fwidjaya20/symphonic-skeleton/shared/context"
	"github.com/fwidjaya20/symphonic-skeleton/src/calendar/application/listeners"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/event/task"
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
			listeners.NewTaskCreatedSubscriber(ctx),
		},
	}
}
