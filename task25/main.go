package main

import (
	"log"
	"time"
)

func sleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	log.Println("start")
	sleep(3 * time.Second)
	log.Println("stop")
}
