package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	w := strings.Fields(s)
	l := len(w)

	for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}

	return strings.Join(w, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverse(s))
}
