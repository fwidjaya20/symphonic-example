package application

import (
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/commands"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/queries"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/service"
)

type TaskApplication struct {
	Create  commands.CreateHandler
	All     queries.AllHandler
	GetById queries.GetByIdHandler
}

func NewTaskApplication(service service.TaskService) TaskApplication {
	return TaskApplication{
		Create:  commands.NewCreateHandler(service),
		All:     queries.NewAllHandler(service),
		GetById: queries.NewGetByIdHandler(service),
	}
}
