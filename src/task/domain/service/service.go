package service

import (
	"github.com/fwidjaya20/symphonic-example/src/task/domain/repository"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/service/task"
)

type TaskService struct {
	Create  task.CreateHandler
	All     task.AllHandler
	GetById task.GetByIdHandler
}

func NewTaskService(repository repository.Repository) TaskService {
	return TaskService{
		Create:  task.NewCreateHandler(repository),
		All:     task.NewAllHandler(repository),
		GetById: task.NewGetByIdHandler(repository),
	}
}
