package storage

import (
	"context"

	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/internal/interfaces"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/memory"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/memslice"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/postgresdb"
)

//NewStorageDB create storage for calendar
func NewStorageDB(cfg *config.Config) (interfaces.EventStorage, error) {
	var err error
	var db interfaces.EventStorage

	switch cfg.DB.DbName {
	case "MemorySlice": // MemorySlice хранение событий в slice
		db = memslice.NewEventStore()
	case "MemoryMap": // MemoryMap хранение событий в map
		db = memory.NewEventStore()

	case "Postgres": //Postgres работа с БД
		ctx := context.Background()
		db, err = postgresdb.NewPgEventStore(ctx, cfg.DB.DatabaseURL)
		if err != nil {
			return nil, err
		}
	}
	return db, err
}

//NewStorageRemind create storage for calendar
func NewStorageRemind(cfg *config.Config) (interfaces.UseCaseReminder, error) {
	var err error
	var db interfaces.UseCaseReminder

	switch cfg.DB.DbName {
	case "MemorySlice": // MemorySlice хранение событий в slice
		db = memslice.NewEventStore()
	case "MemoryMap": // MemoryMap хранение событий в map
		db = memory.NewEventStore()

	case "Postgres": //Postgres работа с БД
		ctx := context.Background()
		db, err = postgresdb.NewPgEventStore(ctx, cfg.DB.DatabaseURL)
		if err != nil {
			return nil, err
		}
	}
	return db, err
}
