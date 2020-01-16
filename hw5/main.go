package main

import (
	log "github.com/sirupsen/logrus"
	rt "hw5/runtask"
)

func main() {
	log.SetLevel(log.InfoLevel)

	countWorker := 10
	ff1 := rt.GenTask(20, 0)                         //генерируем функции для выполнения
	countRunTask, err := rt.Run(ff1, countWorker, 0) // выполняем
	if err != nil {
		log.Error(err)
	} else {
		log.Printf("Ok. Выполнено заданий: %d\n", countRunTask)
	}

	log.Println("====================================================")

	ff2 := rt.GenTask(20, 3)
	countRunTask, err = rt.Run(ff2, countWorker, 3)
	if err != nil {
		log.Error(err)
	} else {
		log.Printf("Ok. Выполнено заданий: %d\n", countRunTask)
	}

}
