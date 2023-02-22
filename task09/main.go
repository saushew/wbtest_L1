package main

import "fmt"

func doubler(ch <-chan int, chDouble chan<- int) {
	for {
		a, ok := <-ch
		if !ok {
			close(chDouble)
			return
		}
		chDouble <- 2 * a
	}
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	ch := make(chan int)
	chDouble := make(chan int)

	go doubler(ch, chDouble)

	go func() {
		for _, a := range arr {
			ch <- a
		}
	}()

	for range arr {
		fmt.Println(<-chDouble)
	}
	close(ch)
}
