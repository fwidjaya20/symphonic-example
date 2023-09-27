package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	Id          uuid.UUID
	Title       string
	Description sql.NullString
	IsCompleted bool
	IsPriority  bool
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (m *Task) TableName() string {
	return "tasks"
}

func (m *Task) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewRandom()
	if nil != err {
		return err
	}

	m.Id = uid

	return nil
}
