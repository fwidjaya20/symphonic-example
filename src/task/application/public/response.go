package public

import (
	"time"

	"github.com/google/uuid"
)

type TaskResponse struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	IsCompleted bool      `json:"is_completed"`
	IsPriority  bool      `json:"is_priority"`
	DueDate     time.Time `json:"due_date"`
	FmtDueDate  string    `json:"fmt_due_date"`
}
