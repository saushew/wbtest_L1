package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func func0() {
	for {
	}
}

func func1(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("func2 completed")

	return
}

func func2(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic in func1:", r)
		}
	}()

	panic("func1 completed")
}

func func3(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("func3 completed")

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func func4(stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("func4 completed")

	for {
		select {
		case <-stop:
			return
		}
	}
}

func func5(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("func4 completed")

	runtime.Goexit()
}

func main() {
	wg := &sync.WaitGroup{}

	// завершение с основным потоком
	go func0()

	// заверешение по выполнению работы
	wg.Add(1)
	go func1(wg)

	// паника
	wg.Add(1)
	go func2(wg)

	// остановка контекстом
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func3(ctx, wg)
	cancel()

	// остановка отдельным stop каналом
	wg.Add(1)
	stop := make(chan struct{})
	go func4(stop, wg)
	close(stop)

	// остановка через runtime.Goexit()
	wg.Add(1)
	go func5(wg)

	wg.Wait()
}
