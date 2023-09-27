package task

import (
	"fmt"

	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/public"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/model"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/repository"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
)

type GetByIdHandler struct {
	repository repository.Repository
}

func (h GetByIdHandler) Execute(c echo.Context, request public.GetTaskRequest) (*model.Task, error) {
	var (
		err  error
		task *model.Task
	)

	if task, err = h.repository.FindOne(c, request.Id); nil != err {
		facades.Logger().Errorf("[Task] get all tasks error: %v", err.Error())
		return nil, fmt.Errorf("failed to retrieve all tasks")
	}

	return task, nil
}

func NewGetByIdHandler(repository repository.Repository) GetByIdHandler {
	return GetByIdHandler{
		repository: repository,
	}
}
