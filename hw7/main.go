package main

import (
	"fmt"
	"os"

	"github.com/mpuzanov/otus-go/hw7/internal/app/goenvdir"
)

func main() {
	fmt.Println(os.Args[:])

	if len(os.Args) < 3 {
		fmt.Println("usage: go-envdir /path/to/evndir command arg1 arg2")
		os.Exit(0)
	}
	//fmt.Println("path:", os.Args[1])
	env, err := goenvdir.ReadDir(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//fmt.Println("command arg1 arg2:", os.Args[2:])
	res := goenvdir.RunCmd(os.Args[2:], env)
	fmt.Println(res)

}
