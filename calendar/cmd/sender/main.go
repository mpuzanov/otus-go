package main

import (
	"log"
)

func main() {
	if err := serverCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
