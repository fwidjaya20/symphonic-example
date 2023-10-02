package ioc

import (
	"sync"

	Task "github.com/fwidjaya20/symphonic-skeleton/src/task/application"
	TaskEvent "github.com/fwidjaya20/symphonic-skeleton/src/task/domain/event"
	TaskService "github.com/fwidjaya20/symphonic-skeleton/src/task/domain/service"
	TaskRepository "github.com/fwidjaya20/symphonic-skeleton/src/task/infrastructure/database"
)

type Container struct {
	Task Task.TaskApplication
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
	taskService := TaskService.NewTaskService(taskRepository)

	return Container{
		Task: Task.NewTaskApplication(taskEvent, taskService),
	}
}
