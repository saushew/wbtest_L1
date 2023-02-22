package main

import (
	"fmt"
	"time"
)

// Adaptee

type oldIntefrace interface {
	printTime() time.Time
}

type service1 struct {
	t time.Time
}

func (s1 *service1) printTime() time.Time {
	return s1.t
}

//~~~~~~~~~~~~~~~

// Target

type newIntefrace interface {
	printTime() int64
}

//~~~~~~~~~~~~~~~

// Adapter

type adapter struct {
	*service1
}

func newAdapter(service1 *service1) newIntefrace {
	return &adapter{service1}
}

func (a *adapter) printTime() int64 {
	return a.t.Unix()
}

func main() {
	s1 := &service1{t: time.Now()}
	s2 := newAdapter(s1)
	
	fmt.Println(s1.printTime())
	fmt.Println(s2.printTime())
}
