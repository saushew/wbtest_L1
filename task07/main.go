package main

import (
	"math/rand"
	"sync"
)

type safeMap struct {
	sync.Mutex
	m map[int]int
}

func safeWriter(sM *safeMap, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		sM.Lock()
		sM.m[1] = rand.Intn(10)
		sM.Unlock()
	}
}

func main() {
	sM := &safeMap{m: make(map[int]int)}

	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go safeWriter(sM, wg)
	}

	wg.Wait()
}
