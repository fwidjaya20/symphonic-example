package listeners

import (
	"github.com/fwidjaya20/symphonic-skeleton/src/calendar/domain/entity"
	"github.com/fwidjaya20/symphonic-skeleton/src/calendar/infrastructure/calendar"
	"github.com/fwidjaya20/symphonic-skeleton/src/calendar/infrastructure/calendar/google"
	TaskEvent "github.com/fwidjaya20/symphonic-skeleton/src/task/domain/event/task"
	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
)

type TaskCreatedSubscriber struct {
	calendar calendar.Calendar
	ctx      echo.Context
}

func (l *TaskCreatedSubscriber) Signature() string {
	return "Task Created Subscriber"
}

func (l *TaskCreatedSubscriber) Handle(job ContractEvent.Job) error {
	task := job.GetPayload().(TaskEvent.TaskCreated)

	if err := l.calendar.CreateEvent(entity.Event{
		Title:       task.Title,
		Description: task.Description,
		IsPriority:  task.IsPriority,
		DueDate:     task.DueDate,
	}); nil != err {
		facades.Logger().Errorf("[Calendar] create new event on calendar error: %v", err.Error())
		return err
	}
	return nil
}

func NewTaskCreatedSubscriber(ctx echo.Context) *TaskCreatedSubscriber {
	return &TaskCreatedSubscriber{
		calendar: google.NewGoogleCalendar(),
		ctx:      ctx,
	}
}
