package commands

import (
	"github.com/fwidjaya20/symphonic-skeleton/shared/exception"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/public"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/constant"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/event"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/model"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/service"
	"github.com/golang-module/carbon"
	"github.com/labstack/echo/v4"
)

type CreateHandler struct {
	event   event.TaskEvent
	service service.TaskService
}

func (h CreateHandler) Execute(c echo.Context, request public.CreateTaskRequest) (*public.TaskResponse, error) {
	var (
		err  error
		task *model.Task
	)

	if task, err = h.service.Create.Execute(c, request); nil != err {
		return nil, exception.New(err, constant.ErrCreateRecord, err.Error(), nil)
	}

	if err = h.event.TaskCreated.Publish(*task); nil != err {
		return nil, exception.New(err, constant.ErrPublishCreatedRecord, err.Error(), nil)
	}

	response := public.TaskResponse{
		Id:          task.Id,
		Title:       task.Title,
		IsCompleted: task.IsCompleted,
		IsPriority:  task.IsPriority,
		DueDate:     task.DueDate,
		FmtDueDate:  carbon.Parse(task.DueDate.String()).Format("M, d Y H:i:s O"),
	}

	if task.Description.Valid {
		response.Description = &task.Description.String
	}

	return &response, nil
}

func NewCreateHandler(event event.TaskEvent, service service.TaskService) CreateHandler {
	return CreateHandler{
		event:   event,
		service: service,
	}
}
