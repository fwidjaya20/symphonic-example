package google

import (
	"fmt"
	"strings"

	"github.com/fwidjaya20/symphonic-example/src/calendar/domain/entity"
	"github.com/fwidjaya20/symphonic-example/src/calendar/infrastructure/calendar"
	"github.com/golang-module/carbon"
)

type GoogleCalendar struct{}

func NewGoogleCalendar() calendar.Calendar {
	return &GoogleCalendar{}
}

func (g *GoogleCalendar) CreateEvent(event entity.Event) error {
	var formattedStr = make([]string, 0)

	if event.IsPriority {
		formattedStr = append(formattedStr, "[PRIORITY]")
	}

	formattedStr = append(formattedStr, event.Title, carbon.Parse(event.DueDate.String()).Format("M, d Y H:i:s O"))

	fmt.Println("[Google Calendar]", strings.Join(formattedStr, " "))

	return nil
}
