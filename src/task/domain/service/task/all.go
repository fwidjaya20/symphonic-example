package task

import (
	"fmt"

	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/model"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/repository"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
)

type AllHandler struct {
	repository repository.Repository
}

func (h AllHandler) Execute(c echo.Context) ([]model.Task, error) {
	var (
		err   error
		tasks []model.Task
	)

	if tasks, err = h.repository.All(c); nil != err {
		facades.Logger().Errorf("[Task] get all tasks error: %v", err.Error())
		return nil, fmt.Errorf("failed to retrieve all tasks")
	}

	return tasks, nil
}

func NewAllHandler(repository repository.Repository) AllHandler {
	return AllHandler{
		repository: repository,
	}
}
