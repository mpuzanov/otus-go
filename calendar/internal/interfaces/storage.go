package interfaces

import (
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//EventStorage интерфейс для работы с DB
type EventStorage interface {
	AddEvent(event *model.Event) (string, error)
	UpdateEvent(event *model.Event) (bool, error)
	DelEvent(id string) (bool, error)
	FindEventByID(id string) (*model.Event, error)
	GetUserEvents(user string) ([]model.Event, error)
}

//UseCaseReminder интерфейс для работы scheduler
type UseCaseReminder interface {
	GetUserEvents(user string) ([]model.Event, error)
}
