package application

import (
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/commands"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/queries"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/event"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/service"
)

type TaskApplication struct {
	Create  commands.CreateHandler
	All     queries.AllHandler
	GetById queries.GetByIdHandler
}

func NewTaskApplication(event event.TaskEvent, service service.TaskService) TaskApplication {
	return TaskApplication{
		Create:  commands.NewCreateHandler(event, service),
		All:     queries.NewAllHandler(service),
		GetById: queries.NewGetByIdHandler(service),
	}
}
