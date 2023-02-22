package main

import "fmt"

func set(arr []string) []string {
	m := make(map[string]struct{})

	for _, v := range arr {
		m[v] = struct{}{}
	}

	result := make([]string, 0, len(m))
	for key := range m {
		result = append(result, key)
	}

	return result
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(set(arr))
}
