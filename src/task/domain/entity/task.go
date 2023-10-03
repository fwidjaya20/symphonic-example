package entity

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-module/carbon"
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

func (t *Task) TableName() string {
	return "tasks"
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewRandom()
	if nil != err {
		return err
	}

	t.Id = uid

	return nil
}

func (t *Task) IsValid() (bool, error) {
	var err error

	if carbon.Parse(t.DueDate.String()).IsPast() {
		err = fmt.Errorf("the 'due_date' on this task was incorrect")
	}

	return nil == err, err
}
