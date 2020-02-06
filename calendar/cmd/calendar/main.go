package main

import (
	"log"

	"github.com/mpuzanov/otus-go/calendar/internal/apiserver"
)

func main() {
	var err error

	if err = apiserver.Start(); err != nil {
		log.Fatal(err)
	}

}
