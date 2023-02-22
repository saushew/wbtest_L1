package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var numOfWorkers string

func init() {
	flag.StringVar(&numOfWorkers, "w", "5", "number of workers")
}

func worker(jobs <-chan int, finish chan<- struct{}, stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			finish <- struct{}{}
			return
		case s := <-jobs:
			log.Println(s)
			time.Sleep(5 * time.Second)
		}
	}
}

func main() {
	flag.Parse()

	log.Printf("%s workers will be started\n", numOfWorkers)

	numOfW, err := strconv.Atoi(numOfWorkers)
	if err != nil {
		log.Println(err)
		return
	}

	// jobs - канал для записи в главном потоке, воркеры из него читают
	jobs := make(chan int)
	// stop - канал чтобы остановить публикации в основной и остановить все воркеры
	stop := make(chan struct{})
	// finish - через этот канал воркеры сообщат, что они остановились
	finish := make(chan struct{})

	// запускаем воркеры
	for w := 0; w < numOfW; w++ {
		go worker(jobs, finish, stop)
	}

	// канал для сигнала остановки ctrl+C
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// публикация в канал
	go func() {
		for i := 0; true; i++ {
			select {
			case <-stop:
				return
			case jobs <- i:
			}
		}
	}()

	// ждем сигнал ctrl+C
	select {
	case s := <-interrupt:
		close(stop)
		log.Println(s.String())
	}

	// ждем сообщения об остановке от воркеров
	for w := 0; w < numOfW; w++ {
		<-finish
	}

	close(jobs)
	close(finish)
}
