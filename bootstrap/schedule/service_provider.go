package schedule

import (
	"github.com/fwidjaya20/symphonic/contracts/foundation"
	"github.com/fwidjaya20/symphonic/facades"
)

type TaskSchedulerServiceProvider struct {
}

func (sp *TaskSchedulerServiceProvider) Boot(app foundation.Application) {}

func (sp *TaskSchedulerServiceProvider) Register(app foundation.Application) {
	kernel := Kernel{}

	facades.Schedule().Register(kernel.Schedule())
}
