package calendar

import "github.com/fwidjaya20/symphonic-skeleton/src/calendar/domain/entity"

type Calendar interface {
	CreateEvent(event entity.Event) error
}
