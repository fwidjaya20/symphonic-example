package ioc

import (
	"sync"

	Notification "github.com/fwidjaya20/symphonic-example/src/notification/application"
	Task "github.com/fwidjaya20/symphonic-example/src/task/application"
	TaskEvent "github.com/fwidjaya20/symphonic-example/src/task/domain/event"
	TaskRepository "github.com/fwidjaya20/symphonic-example/src/task/infrastructure/repository"
)

type Container struct {
	Notification Notification.NotificationApplication
	Task         Task.TaskApplication
}

var (
	instanceIoC Container
	syncIoC     sync.Once
)

func Injector() Container {
	syncIoC.Do(func() {
		instanceIoC = newContainer()
	})

	return instanceIoC
}

func newContainer() Container {
	taskRepository := TaskRepository.NewPostgresRepository()
	taskEvent := TaskEvent.NewTaskEvent()

	return Container{
		Notification: Notification.NewNotificationApplication(),
		Task:         Task.NewTaskApplication(taskEvent, taskRepository),
	}
}
