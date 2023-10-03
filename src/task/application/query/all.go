package query

import (
	"github.com/fwidjaya20/symphonic-example/shared/exception"
	"github.com/fwidjaya20/symphonic-example/src/task/application/public"
	"github.com/fwidjaya20/symphonic-example/src/task/constant"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/entity"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/interface/repository"
	"github.com/fwidjaya20/symphonic-example/src/task/transformer"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
)

type AllHandler struct {
	repository repository.Repository
}

func (h AllHandler) Execute(c echo.Context, request public.GetTasksRequest) ([]public.TaskResponse, error) {
	var (
		err      error
		tasks    []entity.Task
		response = make([]public.TaskResponse, 0)
	)

	if tasks, err = h.repository.All(c, request); nil != err {
		facades.Logger().Errorf("[Task] get all tasks error: %v", err.Error())
		return response, exception.New(err, constant.ErrRetrieveRecord, "failed to retrieve all tasks", nil)
	}

	response = transformer.TransformTaskEntitiesAsResponses(tasks)

	return response, nil
}

func NewAllHandler(repository repository.Repository) AllHandler {
	return AllHandler{
		repository: repository,
	}
}
