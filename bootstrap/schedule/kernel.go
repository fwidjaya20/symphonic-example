package schedule

import (
	SharedContext "github.com/fwidjaya20/symphonic-example/shared/context"
	TaskSchedule "github.com/fwidjaya20/symphonic-example/src/task/application/schedule"
	"github.com/fwidjaya20/symphonic/contracts/schedule"
	"github.com/labstack/echo/v4"
)

type Kernel struct{}

func (kernel *Kernel) Schedule() []schedule.Job {
	e := echo.New()

	ctx := &SharedContext.SymphonicContext{
		Context: e.NewContext(nil, nil),
	}

	return []schedule.Job{
		TaskSchedule.TaskReminder(ctx),
	}
}
