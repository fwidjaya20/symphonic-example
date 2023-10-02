package schedule

import (
	SharedContext "github.com/fwidjaya20/symphonic-skeleton/shared/context"
	TaskSchedule "github.com/fwidjaya20/symphonic-skeleton/src/task/application/schedules"
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
