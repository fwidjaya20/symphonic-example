package application

import (
	"github.com/fwidjaya20/symphonic-example/src/task/application/command"
	"github.com/fwidjaya20/symphonic-example/src/task/application/query"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/event"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/service"
)

type TaskApplication struct {
	Create  command.CreateHandler
	All     query.AllHandler
	GetById query.GetByIdHandler
}

func NewTaskApplication(event event.TaskEvent, service service.TaskService) TaskApplication {
	return TaskApplication{
		Create:  command.NewCreateHandler(event, service),
		All:     query.NewAllHandler(service),
		GetById: query.NewGetByIdHandler(service),
	}
}
