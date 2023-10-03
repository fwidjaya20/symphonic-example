package repository

import (
	"github.com/fwidjaya20/symphonic-example/src/task/application/public"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/entity"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(c echo.Context, m *entity.Task) error

	All(c echo.Context, payload public.GetTasksRequest) ([]entity.Task, error)
	FindOne(c echo.Context, id uuid.UUID) (*entity.Task, error)
}
