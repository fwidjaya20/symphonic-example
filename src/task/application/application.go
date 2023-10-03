package application

import (
	"github.com/fwidjaya20/symphonic-example/src/task/application/command"
	"github.com/fwidjaya20/symphonic-example/src/task/application/query"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/event"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/interface/repository"
)

type TaskApplication struct {
	Create  command.CreateHandler
	All     query.AllHandler
	GetById query.GetByIdHandler
}

func NewTaskApplication(event event.TaskEvent, repository repository.Repository) TaskApplication {
	return TaskApplication{
		Create:  command.NewCreateHandler(event, repository),
		All:     query.NewAllHandler(repository),
		GetById: query.NewGetByIdHandler(repository),
	}
}
