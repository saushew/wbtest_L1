package main

import (
	"fmt"
)

func reverse(s string) string {
	r := []rune(s)
	l := len(r)

	for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func main() {
	s := "главрыба"
	fmt.Println(reverse(s))
}
