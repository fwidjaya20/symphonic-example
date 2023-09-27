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

type AllHandler struct {
	service service.TaskService
}

func (h AllHandler) Execute(c echo.Context) ([]public.TaskResponse, error) {
	var (
		err      error
		tasks    []model.Task
		response []public.TaskResponse
	)

	if tasks, err = h.service.All.Execute(c); nil != err {
		return nil, exception.New(err, constant.ErrCreateRecord, err.Error(), nil)
	}

	for _, it := range tasks {
		obj := public.TaskResponse{
			Id:          it.Id,
			Title:       it.Title,
			IsCompleted: it.IsCompleted,
			IsPriority:  it.IsPriority,
			DueDate:     it.DueDate,
			FmtDueDate:  carbon.Parse(it.DueDate.String()).Format("M, D Y H:i:s"),
		}

		if it.Description.Valid {
			obj.Description = &it.Description.String
		}

		response = append(response, obj)
	}

	return response, nil
}

func NewAllHandler(service service.TaskService) AllHandler {
	return AllHandler{
		service: service,
	}
}
