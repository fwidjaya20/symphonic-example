package command

import (
	"fmt"

	"github.com/fwidjaya20/symphonic-example/shared/exception"
	"github.com/fwidjaya20/symphonic-example/src/task/application/public"
	"github.com/fwidjaya20/symphonic-example/src/task/constant"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/entity"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/event"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/interface/repository"
	"github.com/fwidjaya20/symphonic-example/src/task/transformer"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
)

type CreateHandler struct {
	event      event.TaskEvent
	repository repository.Repository
}

func (h CreateHandler) Execute(c echo.Context, request public.CreateTaskRequest) (*public.TaskResponse, error) {
	var (
		err  error
		task entity.Task
	)

	task = transformer.TransformRequestToEntity(request)

	if err := h.repository.Create(c, &task); nil != err {
		facades.Logger().Errorf("[Task] create new task error: %v", err.Error())
		return nil, fmt.Errorf("failed to create new task")
	}

	if err = h.event.TaskCreated.Publish(task); nil != err {
		return nil, exception.New(err, constant.ErrPublishCreatedRecord, err.Error(), nil)
	}

	response := transformer.TransformTaskEntityAsResponse(task)

	return &response, nil
}

func NewCreateHandler(event event.TaskEvent, repository repository.Repository) CreateHandler {
	return CreateHandler{
		event:      event,
		repository: repository,
	}
}
