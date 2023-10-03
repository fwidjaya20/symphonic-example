package listener

import (
	"fmt"

	"github.com/fwidjaya20/symphonic-example/ioc"
	NotificationEnum "github.com/fwidjaya20/symphonic-example/shared/enum/notification"
	"github.com/fwidjaya20/symphonic-example/src/notification/application/public"
	TaskEvent "github.com/fwidjaya20/symphonic-example/src/task/domain/event/task"
	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/golang-module/carbon"
	"github.com/labstack/echo/v4"
)

type TaskCreatedSubscriber struct {
	ctx echo.Context
}

func (l *TaskCreatedSubscriber) Signature() string {
	return "[Notification] Task Created Subscriber"
}

func (l *TaskCreatedSubscriber) Handle(job ContractEvent.Job) error {
	task := job.GetPayload().(TaskEvent.TaskCreated)

	if err := ioc.Injector().Notification.Send.Execute(l.ctx, public.SendNotificationRequest{
		Channel: NotificationEnum.FirebaseCloudMessaging,
		From:    "no-reply@todo.app",
		To:      "T9Si7vuA62PHg6X3+lTVUB9wxCOk73daX5Fn72dBisw=MOCK",
		Subject: "New Task Created",
		Body:    fmt.Sprintf("A new task '%s' has been added to your to-do list. Please make sure to complete the task before %s.", task.Title, carbon.Parse(task.DueDate.String()).Format("M, d Y H:i:s O")),
	}); nil != err {
		return err
	}

	return nil
}

func NewTaskCreatedSubscriber(ctx echo.Context) *TaskCreatedSubscriber {
	return &TaskCreatedSubscriber{
		ctx: ctx,
	}
}
