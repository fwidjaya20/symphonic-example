package commands

import (
	"github.com/fwidjaya20/symphonic-skeleton/shared/exception"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/public"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/constant"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/model"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/service"
	"github.com/golang-module/carbon"
	"github.com/labstack/echo/v4"
)

type CreateHandler struct {
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

func NewCreateHandler(service service.TaskService) CreateHandler {
	return CreateHandler{
		service: service,
	}
}
