package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var justString string

func createHugeString(n int) string {
	b := strings.Builder{}
	b.Grow(n)

	rand.Seed(time.Now().Unix())

	for i := 0; i < n; i++ {
		r := byte(rand.Intn(26) + 'a')
		b.WriteByte(r)
	}

	return b.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	// justString = v[:100]
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()

	fmt.Println(justString)
}
