package task

import (
	"time"

	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/model"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/google/uuid"
)

type TaskCreated struct {
	Id          uuid.UUID
	Title       string
	Description *string
	IsPriority  bool
	DueDate     time.Time
}

func (e TaskCreated) Signature() string {
	return "Task.Created"
}

func (e TaskCreated) GetPayload() any {
	return e
}

type PublisherTaskCreated struct{}

func (p PublisherTaskCreated) Publish(task model.Task) error {
	event := TaskCreated{
		Id:          task.Id,
		Title:       task.Title,
		Description: nil,
		IsPriority:  task.IsPriority,
		DueDate:     task.DueDate,
	}

	if task.Description.Valid {
		event.Description = &task.Description.String
	}

	if err := facades.Event().Job(&event).Publish(); nil != err {
		facades.Logger().Errorf("[Task] publish task created error: %v", err.Error())
		return err
	}

	return nil
}
