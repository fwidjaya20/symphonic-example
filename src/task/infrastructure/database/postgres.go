package database

import (
	"errors"

	"github.com/fwidjaya20/symphonic-skeleton/shared/context"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/public"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/model"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/domain/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type postgres struct{}

func NewPostgresRepository() repository.Repository {
	return &postgres{}
}

func (p postgres) Create(c echo.Context, m *model.Task) error {
	return c.(*context.SymphonicContext).Database.Create(&m).Error
}

func (p postgres) All(c echo.Context, payload public.GetTasksRequest) ([]model.Task, error) {
	var (
		err     error
		results = make([]model.Task, 0)
	)

	ctx := c.(*context.SymphonicContext)

	if err = ctx.Database.
		Model(&model.Task{}).
		Where(`"tasks"."is_completed" = ?`, payload.IsCompleted).
		Order(`"tasks"."due_date" ASC`).
		Order(`"tasks"."is_priority" DESC`).
		Find(&results).
		Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return results, nil
		}
		return nil, err
	}

	return results, nil
}

func (p postgres) FindOne(c echo.Context, id uuid.UUID) (*model.Task, error) {
	var (
		err    error
		result model.Task
	)

	ctx := c.(*context.SymphonicContext)

	if err = ctx.Database.
		Model(&model.Task{}).
		First(&result).
		Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
