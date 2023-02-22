package main

import (
	"fmt"
	"math"
)

func changeBit(num int64, position int) int64 {
	a := int64(math.Pow(2, float64(position)))
	return num ^ a
}

func main() {
	var num int64
	var position int

	num = 1
	position = 64

	fmt.Println(changeBit(num, position))
}
