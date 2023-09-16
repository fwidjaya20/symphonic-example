package schedule

import (
	"github.com/fwidjaya20/symphonic/contracts/schedule"
)

type Kernel struct{}

func (kernel *Kernel) Schedule() []schedule.Job {
	return []schedule.Job{}
}
