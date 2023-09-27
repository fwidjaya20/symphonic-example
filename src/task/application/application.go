package application

import (
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/commands"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/service"
)

type TaskApplication struct {
	Create commands.CreateHandler
}

func NewTaskApplication(service service.TaskService) TaskApplication {
	return TaskApplication{
		Create: commands.NewCreateHandler(service),
	}
}
