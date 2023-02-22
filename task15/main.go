package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func createHugeString(v *string, n int) {
	b := strings.Builder{}
	b.Grow(n)

	rand.Seed(time.Now().Unix())

	for i := 0; i < n; i++ {
		r := string(rand.Intn(31) + 1072)
		b.WriteString(r)
	}

	*v = b.String()
}

func someFunc(justString *string) {
	var v string

	createHugeString(&v, 1<<10)

	*justString = string([]rune(v)[:100])
}

func main() {
	var justString string

	someFunc(&justString)

	fmt.Println(justString)
}
