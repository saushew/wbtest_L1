package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	array := []int64{2, 4, 6, 8, 10}

// 	var sumOfSquares int64

// 	wg := &sync.WaitGroup{}
// 	for _, val := range array {
// 		wg.Add(1)
// 		go func(a int64) {
// 			atomic.AddInt64(&sumOfSquares, a*a)
// 			wg.Done()
// 		}(val)
// 	}

// 	wg.Wait()

// 	fmt.Println(sumOfSquares)
// }

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func main() {
	array := []int64{2, 4, 6, 8, 10}

	var sumOfSquares struct {
		sync.Mutex
		sum int64
	}

	wg := &sync.WaitGroup{}
	for _, val := range array {
		wg.Add(1)
		go func(a int64) {
			sumOfSquares.Lock()
			sumOfSquares.sum += a*a
			sumOfSquares.Unlock()

			wg.Done()
		}(val)
	}

	wg.Wait()

	fmt.Println(sumOfSquares.sum)
}
