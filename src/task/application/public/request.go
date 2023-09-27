package public

import (
	"time"

	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	DueDate     time.Time `json:"due_date"`
	IsPriority  bool      `json:"is_priority"`
}

type GetTaskRequest struct {
	Id uuid.UUID `param:"task_id"`
}
