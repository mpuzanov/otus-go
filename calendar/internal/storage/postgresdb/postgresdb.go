package postgresdb

import (
	"context"
	"database/sql"

	// Register some standard stuff
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mpuzanov/otus-go/calendar/internal/errors"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//EventStore структура хранения списка событий Календаря
type EventStore struct {
	ctx context.Context
	db  *sqlx.DB
}

//NewPgEventStore Возвращаем хранилище
func NewPgEventStore(ctx context.Context, databaseURL string) (*EventStore, error) {
	db, err := sqlx.Open("pgx", databaseURL) // *sql.DB
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &EventStore{ctx: ctx, db: db}, nil
}

//AddEvent Добавляем событие
func (pg *EventStore) AddEvent(event *model.Event) (string, error) {
	eDb, err := toEventDB(event)
	if err != nil {
		return "", err
	}
	query := `INSERT INTO events (header, text, start_time, end_time, user_id, reminder_before) VALUES 
	(:Header, :Text, :StartTime, :EndTime, :UserName, :ReminderBefore) RETURNING id`
	var id string
	p := map[string]interface{}{
		"Header":         eDb.Header,
		"Text":           eDb.Text,
		"StartTime":      eDb.StartTime,
		"EndTime":        eDb.EndTime,
		"UserName":       eDb.UserName,
		"ReminderBefore": eDb.ReminderBefore.Get(),
	}
	nstmt, err := pg.db.PrepareNamedContext(pg.ctx, query)
	if err != nil {
		return "", err
	}
	defer nstmt.Close()
	if err := nstmt.GetContext(pg.ctx, &id, p); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.ErrNoDBAffected
		}
		return "", err
	}
	return id, nil
}

//UpdateEvent Изменение события
func (pg *EventStore) UpdateEvent(event *model.Event) (bool, error) {
	eDb, err := toEventDB(event)
	if err != nil {
		return false, err
	}
	p := map[string]interface{}{
		"ID":             eDb.ID,
		"Header":         eDb.Header,
		"Text":           eDb.Text,
		"StartTime":      eDb.StartTime,
		"EndTime":        eDb.EndTime,
		"UserName":       eDb.UserName,
		"ReminderBefore": eDb.ReminderBefore.Get(),
	}
	query := `UPDATE events SET (header, text, start_time, end_time, user_id, reminder_before) = 
	(:Header, :Text, :StartTime, :EndTime, :UserName, :ReminderBefore) WHERE ID = :ID`
	nstmt, err := pg.db.PrepareNamedContext(pg.ctx, query)
	if err != nil {
		return false, err
	}
	defer nstmt.Close()
	res, err := nstmt.ExecContext(pg.ctx, p)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, errors.ErrNoDBAffected
	}
	return true, nil
}

//DelEvent Удаляем событие по коду
func (pg *EventStore) DelEvent(id string) (bool, error) {

	p := map[string]interface{}{"ID": id}
	nstmt, _ := pg.db.PrepareNamedContext(pg.ctx, "DELETE FROM events WHERE id = :ID")
	res, err := nstmt.ExecContext(pg.ctx, p)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, errors.ErrNoDBAffected
	}
	nstmt.Close()
	return true, nil
}

//FindEventByID возврат события по коду
func (pg *EventStore) FindEventByID(id string) (*model.Event, error) {
	evDB := &EventDB{}
	p := map[string]interface{}{"ID": id}
	nstmt, err := pg.db.PrepareNamedContext(pg.ctx, "SELECT * FROM events WHERE id= :ID")
	if err := nstmt.GetContext(pg.ctx, evDB, p); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}
	nstmt.Close()
	ev, err := toEvent(evDB)
	if err != nil {
		return nil, err
	}
	return ev, nil
}

//GetUserEvents получаем события пользователя
func (pg *EventStore) GetUserEvents(user string) ([]model.Event, error) {
	p := map[string]interface{}{"user": user}
	outDB := make([]EventDB, 0)
	out := make([]model.Event, 0)
	nstmt, err := pg.db.PrepareNamedContext(pg.ctx, `select * from events where user_id=:user`)
	if err != nil {
		return nil, err
	}
	defer nstmt.Close()
	if err = nstmt.Select(&outDB, p); err != nil {
		return nil, err
	}
	for _, v := range outDB {
		e, err := toEvent(&v)
		if err != nil {
			return nil, err
		}
		out = append(out, *e)
	}
	return out, nil
}
