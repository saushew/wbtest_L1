package main

import (
	"fmt"
)

func changeBit(num int64, position int) int64 {
	a := int64(1) << position
	return num ^ a
}

func main() {
	var num int64
	var position int

	num = 1
	position = 1

	fmt.Println(changeBit(num, position))
}
