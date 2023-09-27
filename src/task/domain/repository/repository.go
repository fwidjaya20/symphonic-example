package repository

import (
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Create(c echo.Context, m *model.Task) error

	All(c echo.Context) ([]model.Task, error)
	FindOne(c echo.Context, id uuid.UUID) (*model.Task, error)
}
