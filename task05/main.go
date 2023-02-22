package main

import (
	"log"
	"time"
)

func reader(ch <-chan int) {
	for {
		a, ok := <-ch
		if !ok {
			return
		}
		log.Println(a)
	}
}

func main() {
	ch := make(chan int)

	go reader(ch)

	stopT := time.NewTicker(5 * time.Second)
	periodT := time.NewTicker(time.Second)

LOOP:
	for i := 0; true; i++ {
		select {
		case <-stopT.C:
			close(ch)
			break LOOP
		case <-periodT.C:
			ch <- i
		}
	}
}
