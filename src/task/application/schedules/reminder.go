package schedules

import (
	"fmt"

	"github.com/fwidjaya20/symphonic-skeleton/ioc"
	"github.com/fwidjaya20/symphonic-skeleton/shared/context"
	"github.com/fwidjaya20/symphonic-skeleton/shared/database"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/public"
	ContractSchedule "github.com/fwidjaya20/symphonic/contracts/schedule"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/fwidjaya20/symphonic/schedule"
	"github.com/golang-module/carbon"
	"github.com/labstack/echo/v4"
)

func TaskReminder(ctx echo.Context) ContractSchedule.Job {
	return schedule.NewJob(func() {
		var (
			err   error
			tasks []public.TaskResponse
		)

		err = database.RunInTransaction(ctx.(*context.SymphonicContext), func(e echo.Context) error {
			tasks, err = ioc.Injector().Task.All.Execute(ctx, public.GetTasksRequest{IsCompleted: false})
			if nil != err {
				return err
			}
			if len(tasks) == 0 {
				return nil
			}

			for _, task := range tasks {
				if carbon.Parse(task.DueDate.String()).IsTomorrow() {
					fmt.Printf("[Send Notification] Please make sure to complete this task '%s' by tomorrow to stay on track with your schedule.\n", task.Title)
					continue
				}

				if carbon.Parse(task.DueDate.String()).IsToday() {
					fmt.Printf("[Send Notification] We wanted to remind you about an important task '%s' that is due today.\n", task.Title)
					continue
				}
			}

			return err
		})
		if nil != err {
			facades.Logger().Errorf("[Schedule][Task] task reminder error: %v", err.Error())
		}
	}).Daily()
}
