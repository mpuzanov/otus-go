package main

import (
	"fmt"
	"hw2/unpack"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Задайте аргумент для распаковки строки")
		os.Exit(0)
	}
	for _, s := range os.Args[1:] {
		fmt.Println(unpack.Unpack(s))
	}

}
