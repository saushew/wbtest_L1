package main

import (
	"fmt"
)

// func square(a int, wg *sync.WaitGroup) {
// 	fmt.Printf("%d^2 = %d\n", a, a*a)
// 	wg.Done()
// }

// func main() {
// 	array := []int{2, 4, 6, 8, 10}

// 	wg := &sync.WaitGroup{}
// 	for _, val := range array {
// 		wg.Add(1)
// 		go square(val, wg)
// 	}

// 	wg.Wait()
// }

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// func main() {
// 	array := []int{2, 4, 6, 8, 10}

// 	wg := &sync.WaitGroup{}
// 	for _, val := range array {
// 		wg.Add(1)
// 		go func(a int) {
// 			fmt.Printf("%d^2 = %d\n", a, a*a)
// 			wg.Done()
// 		}(val)
// 	}

// 	wg.Wait()
// }

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func square(a int, done chan struct{}) {
	fmt.Printf("%d^2 = %d\n", a, a*a)
	done <- struct{}{}
}

func main() {
	array := []int{2, 4, 6, 8, 10}

	done := make(chan struct{})

	for _, val := range array {
		go square(val, done)
	}

	for range array {
		<-done
	}

	close(done)
}
