package repository

import (
	"errors"

	"github.com/fwidjaya20/symphonic-example/shared/context"
	"github.com/fwidjaya20/symphonic-example/src/task/application/public"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/entity"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/interface/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type postgres struct{}

func NewPostgresRepository() repository.Repository {
	return &postgres{}
}

func (p postgres) Create(c echo.Context, m *entity.Task) error {
	return c.(*context.SymphonicContext).Database.Create(&m).Error
}

func (p postgres) All(c echo.Context, payload public.GetTasksRequest) ([]entity.Task, error) {
	var (
		err     error
		results = make([]entity.Task, 0)
	)

	ctx := c.(*context.SymphonicContext)

	if err = ctx.Database.
		Model(&entity.Task{}).
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

func (p postgres) FindOne(c echo.Context, id uuid.UUID) (*entity.Task, error) {
	var (
		err    error
		result entity.Task
	)

	ctx := c.(*context.SymphonicContext)

	if err = ctx.Database.
		Model(&entity.Task{}).
		First(&result).
		Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
