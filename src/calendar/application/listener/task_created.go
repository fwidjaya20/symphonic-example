package listener

import (
	"github.com/fwidjaya20/symphonic-example/src/calendar/domain/entity"
	DomainInterface "github.com/fwidjaya20/symphonic-example/src/calendar/domain/interface/calendar"
	"github.com/fwidjaya20/symphonic-example/src/calendar/infrastructure/calendar"
	TaskEvent "github.com/fwidjaya20/symphonic-example/src/task/domain/event/task"
	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
)

type TaskCreatedSubscriber struct {
	calendar DomainInterface.Calendar
	ctx      echo.Context
}

func (l *TaskCreatedSubscriber) Signature() string {
	return "[Calendar] Task Created Subscriber"
}

func (l *TaskCreatedSubscriber) Handle(job ContractEvent.Job) error {
	task := job.GetPayload().(TaskEvent.TaskCreated)

	event := entity.Event{
		Title:       task.Title,
		Description: task.Description,
		IsPriority:  task.IsPriority,
		DueDate:     task.DueDate,
	}

	if ok, err := event.IsValid(); !ok {
		facades.Logger().Errorf("[Calendar] incorrect event given: %v", err.Error())
		return err
	}

	if err := l.calendar.CreateEvent(event); nil != err {
		facades.Logger().Errorf("[Calendar] create new event on calendar error: %v", err.Error())
		return err
	}
	return nil
}

func NewTaskCreatedSubscriber(ctx echo.Context) *TaskCreatedSubscriber {
	return &TaskCreatedSubscriber{
		calendar: calendar.NewGoogleCalendar(),
		ctx:      ctx,
	}
}
