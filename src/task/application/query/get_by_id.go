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

type GetByIdHandler struct {
	repository repository.Repository
}

func (h GetByIdHandler) Execute(c echo.Context, request public.GetTaskRequest) (*public.TaskResponse, error) {
	var (
		err      error
		task     *entity.Task
		response public.TaskResponse
	)

	if task, err = h.repository.FindOne(c, request.Id); nil != err {
		facades.Logger().Errorf("[Task] get all tasks error: %v", err.Error())
		return nil, exception.New(err, constant.ErrRetrieveRecord, "failed to retrieve all tasks", nil)
	}

	response = transformer.TransformTaskEntityAsResponse(*task)

	return &response, nil
}

func NewGetByIdHandler(repository repository.Repository) GetByIdHandler {
	return GetByIdHandler{
		repository: repository,
	}
}
