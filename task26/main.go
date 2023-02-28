package main

import (
	"fmt"
	"unicode"
)

func isUnique(s string) bool {
	m := make(map[rune]struct{}, len(s))

	for _, v := range s {
		if !unicode.IsLower(v) {
			v = unicode.ToLower(v)
		}
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}

	return true
}

func main() {
	var s string = "abcCd"
	fmt.Println(isUnique(s))

	s = "abCdefAaf"
	fmt.Println(isUnique(s))

	s = "aabcd"
	fmt.Println(isUnique(s))
}
