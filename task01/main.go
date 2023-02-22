package main

import (
	"fmt"
	"time"
)

// Human .
type Human struct {
	FirstName string
	LastName  string
	BirthDate time.Time
}

// NewHuman .
func NewHuman(firstName, lastName string, birthDate time.Time) *Human {
	return &Human{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
	}
}

// Walk .
func (a *Human) Walk() {
	fmt.Printf("Now %s %s is walking\n", a.FirstName, a.LastName)
}

// Action .
type Action struct {
	*Human
}

func main() {
	person := NewHuman("James", "Smith", time.Date(1990, 2, 14, 0, 0, 0, 0, time.Local))
	actions := Action{person}
	actions.Walk()
}
