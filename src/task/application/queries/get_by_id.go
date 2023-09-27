package queries

import (
	"github.com/fwidjaya20/symphonic-skeleton/shared/exception"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/public"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/constant"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/model"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/service"
	"github.com/golang-module/carbon"
	"github.com/labstack/echo/v4"
)

type GetByIdHandler struct {
	service service.TaskService
}

func (h GetByIdHandler) Execute(c echo.Context, request public.GetTaskRequest) (*public.TaskResponse, error) {
	var (
		err      error
		task     *model.Task
		response public.TaskResponse
	)

	if task, err = h.service.GetById.Execute(c, request); nil != err {
		return nil, exception.New(err, constant.ErrCreateRecord, err.Error(), nil)
	}

	response = public.TaskResponse{
		Id:          task.Id,
		Title:       task.Title,
		IsCompleted: task.IsCompleted,
		IsPriority:  task.IsPriority,
		DueDate:     task.DueDate,
		FmtDueDate:  carbon.Parse(task.DueDate.String()).Format("M, D Y H:i:s"),
	}

	if task.Description.Valid {
		response.Description = &task.Description.String
	}

	return &response, nil
}

func NewGetByIdHandler(service service.TaskService) GetByIdHandler {
	return GetByIdHandler{
		service: service,
	}
}
