package main

import (
	"fmt"
	"log"
	"time"

	m "github.com/mpuzanov/otus-go/calendar/internal/store/memstore"
)

func main() {
	var err error

	fmt.Println("создаём структуру хранения событий")
	listEvent := m.NewEventStore()

	fmt.Println("Добавляем событие")
	event := m.NewEvent("Событие 1", time.Now())
	err = listEvent.AddEvent(event)
	if err != nil {
		log.Fatal(err)
	}
	err = listEvent.AddEvent(m.NewEvent("Событие 2", time.Now()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Показываем результат:")
	fmt.Println(listEvent)

	fmt.Println("Удаляем событие c UUID:", event.UUID)
	err = listEvent.DelEvent(*event)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Показываем результат:")
	fmt.Println(listEvent)
}
