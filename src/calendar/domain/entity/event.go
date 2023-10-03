package entity

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon"
)

type Event struct {
	Title       string
	Description *string
	IsPriority  bool
	DueDate     time.Time
}

func (e Event) IsValid() (bool, error) {
	var (
		err error
	)

	if carbon.Parse(e.DueDate.String()).IsPast() {
		err = fmt.Errorf("the 'due_date' on this event was incorrect")
	}

	return nil == err, err
}
