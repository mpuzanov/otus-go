package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v", err)
		os.Exit(1)
	}
	t := time.Now()
	fmt.Println("Время сервера NTP:            ", response.Time.In(t.Location()))
	fmt.Println("Время компьютера:             ", t)
	fmt.Println("Oтличие от NTP:               ", response.ClockOffset)
	fmt.Println("Время компьютера с учётом NTP:", t.Add(response.ClockOffset))
}
