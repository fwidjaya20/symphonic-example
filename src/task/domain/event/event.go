package event

import "github.com/fwidjaya20/symphonic-skeleton/src/task/domain/event/task"

type TaskEvent struct {
	TaskCreated task.PublisherTaskCreated
}

func NewTaskEvent() TaskEvent {
	return TaskEvent{
		TaskCreated: task.PublisherTaskCreated{},
	}
}
