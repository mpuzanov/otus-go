package memory

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mpuzanov/otus-go/calendar/internal/errors"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//EventStore структура хранения списка событий
type EventStore struct {
	db map[string]model.Event
}

//NewEventStore Возвращаем новое хранилище
func NewEventStore() *EventStore {
	return &EventStore{db: make(map[string]model.Event)}
}

//GetEvents Листинг событий
func (s *EventStore) GetEvents() []model.Event {
	var out []model.Event
	for _, val := range s.db {
		out = append(out, val)
	}
	return out
}

//FindEventByID Найти событие
func (s *EventStore) FindEventByID(id string) (*model.Event, error) {

	ev, exist := s.db[id]
	if exist {
		return &ev, nil
	}
	return nil, errors.ErrNotEvent
}

//AddEvent Добавить событие
func (s *EventStore) AddEvent(event *model.Event) error {
	id := event.ID.String()
	_, exist := s.db[id]
	if exist {
		return fmt.Errorf("событие: %s(%s) уже существует. %w", event.Header, event.ID, errors.ErrAddEvent)
	}
	s.db[id] = *event
	return nil
}

//UpdateEvent Изменить событие
func (s *EventStore) UpdateEvent(event *model.Event) error {
	id := event.ID.String()
	_, exist := s.db[id]
	if !exist {
		return fmt.Errorf("нет события: %s(%s). %w", event.Header, event.ID, errors.ErrEditEvent)
	}
	s.db[id] = *event
	return nil
}

//DelEvent Удалить событие
func (s *EventStore) DelEvent(event *model.Event) error {
	id := event.ID.String()
	_, exist := s.db[id]
	if !exist {
		return fmt.Errorf("нет события: %s(%s). %w", event.Header, event.ID, errors.ErrDelEvent)
	}
	delete(s.db, id)
	return nil
}

//String Печать списка событий
func (s *EventStore) String() {
	for _, ev := range s.db {
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
