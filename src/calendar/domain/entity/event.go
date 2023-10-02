package entity

import "time"

type Event struct {
	Title       string
	Description *string
	IsPriority  bool
	DueDate     time.Time
}
