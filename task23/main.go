package main

import "fmt"

// func deleteIndex(arr []int, index int) []int {
// 	l := len(arr) - 1
// 	arr[index], arr[l] = arr[l], arr[index]
// 	return arr[:l]
// }

func deleteIndex(arr []int, index int) []int {
	result := make([]int, index, len(arr)-1)
	copy(result, arr[:index])
	return append(result, arr[index+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	var index int = 0

	fmt.Println(deleteIndex(arr, index))
}
