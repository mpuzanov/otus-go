package runtask

import (
	"errors"
	"sync"

	log "github.com/sirupsen/logrus"
)

//структура для канала хранения функции
type job func() error

//структура для канала хранения результата выполнения функции
type result struct {
	err error
}

func init() {
	log.SetLevel(log.TraceLevel)
	//log.SetLevel(log.InfoLevel)
}

// Run Выполнение функций из слайса, N - кол-во одновременных обработчиков, M - кол-во ошибок для прерывания выполнения функции
func Run(task []func() error, N int, M int) (int, error) {

	countWorker := N // кол-во обработчиков
	MaxError := M    // кол-во допустимых ошибок
	if MaxError == 0 {
		MaxError = 1
	}
	var countTask int  // кол-во выполенных заданий
	var countError int // кол-во выполенных заданий с ошибками

	log.Infof("Кол-во функций для выполнения: %d. Кол-во обработчиков: %d. Кол-во ошибок для прекращения работы: %d\n", len(task), countWorker, MaxError)

	chanFuncs := make(chan job)          // канал функций для выполнения
	chanResults := make(chan result)     // канал получения результатов
	chanStopFuncs := make(chan struct{}) // канал для прерывания выполнения горутин

	wgFuncs := sync.WaitGroup{}

	// запускаем обработчики
	for i := 0; i < countWorker; i++ {
		wgFuncs.Add(1)
		go func(i int) {
			defer wgFuncs.Done()
			worker(i, chanFuncs, chanResults, chanStopFuncs)
		}(i)
	}

	// ждём окончания всех обработчиков
	go func() {
		defer close(chanResults)
		wgFuncs.Wait()
	}()

	// отправляем функции на выполнение
	go func() {
		defer close(chanFuncs)
		for i := 0; i < len(task); i++ {
			chanFuncs <- task[i]
		}
	}()

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)
	// получаем результаты
	go func() {
		defer wgReceivers.Done()
		// признак закрытия канала для остановки обработчиков
		isStopFunc := false
		// Читаем из канала результатов (подчёт выполнения)
		for r := range chanResults {
			countTask++
			if r.err != nil {
				countError++
			}
			if countError >= MaxError && !isStopFunc {
				log.Tracef("Ошибок: %d. Закрываем канал для отмены выполнения обработчиков\n", countError)
				close(chanStopFuncs)
				isStopFunc = true
			}
		}
	}()
	wgReceivers.Wait()
	log.Tracef("Выполнено заданий: %d, с ошибками: %d\n", countTask, countError)

	return countTask, nil
}

// worker Обработчик для выполнения функций из канала jobs
// res - канал для результатов
// stopJob - канал для прекращения обработки.
func worker(id int, jobs <-chan job, res chan<- result, stopJob <-chan struct{}) {
	for {
		select {
		case <-stopJob:
			log.Tracef("Останавливаем обработчик %d\n", id)
			return
		default:
		}

		select {
		case <-stopJob:
			log.Tracef("Останавливаем обработчик %d\n", id)
			return
		case f, ok := <-jobs:
			if !ok {
				log.Tracef("Заданий больше нет - Останавливаем обработчик %d\n", id)
				return
			}
			var r result
			r.err = f() // выполняем функцию
			res <- r
			log.Tracef("обработчик %d выполнил задание! Ошибка: %v\n", id, r.err)
		}

	}
}

//newTask создание функции для тестов
func newTask(num int, isError bool) func() error {
	if !isError {
		return func() error {
			//log.Trace(num, "task without error")
			return nil
		}
	}
	return func() error {
		//log.Trace(num, "task with an error")
		return errors.New("errors task")
	}
}

// GenTask Генерация среза функций
func GenTask(N int, M int) []func() error {
	log.Infof("GenTask. Count funcs: %d. Errors: %d", N, M)
	var ff []func() error
	for i := 0; i < N; i++ {
		if M > 0 {
			ff = append(ff, newTask(i, true))
			M--
		} else {
			ff = append(ff, newTask(i, false))
		}
	}
	return ff
}
