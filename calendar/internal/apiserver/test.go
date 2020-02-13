package apiserver

import (
	"fmt"
	"log"
	"time"

	c "github.com/mpuzanov/otus-go/calendar/internal/calendar"
)

//sampleWorkCalendar пример работы с календарём
func sampleWorkCalendar() {
	var err error

	fmt.Println("================================================")
	fmt.Println("Проверяем работу календаря")
	fmt.Println("создаём структуру хранения событий")
	err = c.NewCalendar(c.MemorySlice, "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Добавляем 2 события")
	startTime, _ := time.Parse("2006-01-02 15:04", "2020-04-01 09:00")
	endTime, _ := time.Parse("2006-01-02 15:04", "2020-04-01 10:30")

	event1, err := c.DB.CreateEvent("user1", "Событие 1", "", startTime, endTime)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.DB.CreateEvent("user1", "Событие 2", "", startTime, endTime)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Показываем результат:")
	fmt.Println(c.DB.GetEvents())

	fmt.Println("Удаляем событие c UUID:", event1.ID)
	err = c.DB.DelEvent(event1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Показываем результат:")
	fmt.Println(c.DB.GetEvents())
	fmt.Println("================================================")
}
