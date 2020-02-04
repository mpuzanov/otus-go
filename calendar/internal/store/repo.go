package store

import (
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//IEvents интерфейс для работы со структурой календаря
type IEvents interface {
	AddEvent(event model.Event) error
	SetEvent(event model.Event) error
	DelEvent(event model.Event) error
	FindEventByHeader(header string) *model.Event
}
