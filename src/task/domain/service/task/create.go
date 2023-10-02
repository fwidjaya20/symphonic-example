package task

import (
	"database/sql"
	"fmt"

	"github.com/fwidjaya20/symphonic-example/src/task/application/public"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/model"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/repository"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
)

type CreateHandler struct {
	repository repository.Repository
}

func (h CreateHandler) Execute(c echo.Context, request public.CreateTaskRequest) (*model.Task, error) {
	taskModel := model.Task{
		Title:       request.Title,
		Description: sql.NullString{},
		IsCompleted: false,
		IsPriority:  request.IsPriority,
		DueDate:     request.DueDate,
	}

	if request.Description != nil {
		taskModel.Description = sql.NullString{
			String: *request.Description,
			Valid:  true,
		}
	}

	if err := h.repository.Create(c, &taskModel); nil != err {
		facades.Logger().Errorf("[Task] create new task error: %v", err.Error())
		return nil, fmt.Errorf("failed to create new task")
	}

	return &taskModel, nil
}

func NewCreateHandler(repository repository.Repository) CreateHandler {
	return CreateHandler{
		repository: repository,
	}
}
