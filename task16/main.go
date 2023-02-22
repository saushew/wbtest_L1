package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{10, 43, 55, 2, 6, 1}
	sort.Ints(arr)
	fmt.Println(arr)
}
