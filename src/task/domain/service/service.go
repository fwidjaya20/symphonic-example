package service

import (
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/repository"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/service/task"
)

type TaskService struct {
	Create task.CreateHandler
}

func NewTaskService(repository repository.Repository) TaskService {
	return TaskService{
		Create: task.NewCreateHandler(repository),
	}
}
