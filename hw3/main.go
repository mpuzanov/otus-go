package main

import (
	"fmt"
	"hw3/topwords"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Задайте файл для обработки")
		os.Exit(1)
	}
	file := os.Args[1]
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка открытия файла: %s", file)
		os.Exit(1)
	}
	sl := topwords.TopWords10(string(b))
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}
