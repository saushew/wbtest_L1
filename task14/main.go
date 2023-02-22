package main

import "fmt"

func whatIsIt(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	default:
		fmt.Println("I don't know")
	}
}

// func whatIsIt(i interface{}) {
// 	fmt.Printf("%T\n", i)
// }

func main() {
	whatIsIt(100)
	whatIsIt("string")
	whatIsIt(true)
	whatIsIt(make(chan int))
	whatIsIt(make(chan string))

	whatIsIt([3]int{1, 2, 3})
}
