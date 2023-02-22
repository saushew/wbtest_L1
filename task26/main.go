package main

import (
	"fmt"
)

func isUniq(s string) bool {
	m := make(map[rune]struct{}, len(s))

	for _, v := range s {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}

	return true
}

func main() {
	var s string = "abcd"
	fmt.Println(isUniq(s))

	s = "abCdefAaf"
	fmt.Println(isUniq(s))

	s = "aabcd"
	fmt.Println(isUniq(s))
}
