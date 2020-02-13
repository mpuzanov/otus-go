package calendar

import (
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/model"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/memory"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/memslice"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/postgresdb"
)

const (
	//Postgres работа с БД
	Postgres int = iota
	// MemorySlice хранение событий в slice
	MemorySlice
	// MemoryMap хранение событий в map
	MemoryMap
)

//DB интерфейс для взаимодействия с данными на нескольких уровнях
var DB Calendar

//Calendar интерфейс для работы со структурой календаря
type Calendar interface {
	CreateEvent(user, header, text string, startTime time.Time, endTime time.Time) (*model.Event, error)
	AddEvent(event *model.Event) error
	UpdateEvent(event *model.Event) error
	DelEvent(event *model.Event) error
	FindEventByID(id string) (*model.Event, error)
	GetEvents() []model.Event
}

//NewCalendar создание интерфейс для работы со структурой календаря
func NewCalendar(storageType int, dsn string) error {
	var err error

	switch storageType {
	case MemorySlice:
		DB = new(memslice.EventStore)
	case MemoryMap:
		DB = new(memory.EventStore)

	case Postgres:
		DB, err = postgresdb.NewPgEventStore(dsn)
		if err != nil {
			return err
		}
	}
	return nil
}
