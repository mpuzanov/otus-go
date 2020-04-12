package storage

import (
	"context"
	"fmt"

	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/internal/errors"
	"github.com/mpuzanov/otus-go/calendar/internal/interfaces"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/memory"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/memslice"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/postgresdb"
)

//NewStorageDB create storage for calendar
func NewStorageDB(cfg config.DBConf) (interfaces.EventStorage, error) {
	var err error
	var db interfaces.EventStorage

	switch cfg.Name {
	case "MemorySlice": // MemorySlice хранение событий в slice
		db = memslice.NewEventStore()
	case "MemoryMap": // MemoryMap хранение событий в map
		db = memory.NewEventStore()

	case "Postgres": //Postgres работа с БД
		ctx := context.Background()

		if cfg.Name == "" || cfg.User == "" || cfg.Host == "" || cfg.Password == "" || cfg.Database == "" {
			return nil, errors.ErrBadLoginDBConfiguration
		}

		DatabaseURL := fmt.Sprintf("Postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSL)

		db, err = postgresdb.NewPgEventStore(ctx, DatabaseURL)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Connecting to %s:%s/%s the database successfully", cfg.Host, cfg.Port, cfg.Database)
	}
	return db, err
}

//NewStorageRemind create storage for calendar
func NewStorageRemind(cfg config.DBConf) (interfaces.UseCaseReminder, error) {
	var err error
	var db interfaces.UseCaseReminder

	switch cfg.Name {
	case "MemorySlice": // MemorySlice хранение событий в slice
		db = memslice.NewEventStore()
	case "MemoryMap": // MemoryMap хранение событий в map
		db = memory.NewEventStore()

	case "Postgres": //Postgres работа с БД
		ctx := context.Background()

		if cfg.Name == "" || cfg.User == "" || cfg.Host == "" || cfg.Password == "" || cfg.Database == "" {
			return nil, errors.ErrBadLoginDBConfiguration
		}

		DatabaseURL := fmt.Sprintf("Postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSL)
		db, err = postgresdb.NewPgEventStore(ctx, DatabaseURL)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Connecting to %s:%s/%s the database successfully", cfg.Host, cfg.Port, cfg.Database)
	}
	return db, err
}
