package main

import "fmt"

func intersection(a, b []int) []int {
	m := make(map[int]int)

	for _, v := range a {
		m[v]++
	}

	var result []int

	for _, v := range b {
		if i, ok := m[v]; ok && i > 0 {
			m[v]--
			result = append(result, v)
		}
	}

	return result
}

func main() {
	a := []int{1, 1, 2, 3, 4, 5}
	b := []int{6, 3, 5, 2, 1, 1}

	if len(a) < len(b) {
		fmt.Println(intersection(a, b))
	} else {
		fmt.Println(intersection(b, a))
	}
}
