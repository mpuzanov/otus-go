package memslice

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//EventStore структура хранения списка событий
type EventStore struct {
	Events []model.Event
}

//NewEventStore Возвращаем новое хранилище
func NewEventStore() *EventStore {
	return &EventStore{Events: make([]model.Event, 0)}
}

//GetEvents Листинг событий
func (s *EventStore) GetEvents() []model.Event {
	return s.Events
}

//FindEventByID Найти событие
func (s *EventStore) FindEventByID(id string) (*model.Event, error) {
	for _, event := range s.Events {
		if event.ID.String() == id {
			return &event, nil
		}
	}
	return nil, model.ErrNotEvent
}

//AddEvent Добавить событие
func (s *EventStore) AddEvent(event *model.Event) error {
	ev, _ := s.FindEventByID(event.ID.String())
	if ev != nil {
		return fmt.Errorf("событие: %s(%s) уже существует. %w", event.Header, event.ID, model.ErrAddEvent)
	}
	s.Events = append(s.Events, *event)

	return nil
}

//UpdateEvent Изменить событие
func (s *EventStore) UpdateEvent(event *model.Event) error {
	for i, ev := range s.Events {
		if ev.ID == event.ID {
			s.Events[i] = *event
			return nil
		}
	}
	return fmt.Errorf("нет события: %s(%s). %w", event.Header, event.ID, model.ErrEditEvent)
}

//DelEvent Удалить событие
func (s *EventStore) DelEvent(event *model.Event) error {

	for i, ev := range s.Events {
		if ev.ID == event.ID {
			s.Events = append(s.Events[:i], s.Events[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("нет события с кодом: %s. %w", event.ID, model.ErrDelEvent)
}

//String Печать списка событий
func (s *EventStore) String() {
	for _, ev := range s.Events {
		fmt.Println(ev)
	}
}

//CreateEvent Создание и запись события
func (s *EventStore) CreateEvent(user, header, text string, startTime time.Time, endTime time.Time) (*model.Event, error) {
	event := &model.Event{
		ID:        uuid.New(),
		User:      user,
		Header:    header,
		Text:      text,
		StartTime: startTime,
		EndTime:   endTime,
	}
	err := s.AddEvent(event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

//NewEvent Формируем новое событие
func NewEvent(user, header, text string, startTime time.Time, endTime time.Time) *model.Event {
	return &model.Event{
		ID:        uuid.New(),
		User:      user,
		Header:    header,
		Text:      text,
		StartTime: startTime,
		EndTime:   endTime,
	}
}
