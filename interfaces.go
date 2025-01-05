package main

import "fmt"

func main() {
	var x interface{}
	x = "Hello, World"
	fmt.Println(x)
	x = 123
	fmt.Println(x)
	x = true
	fmt.Println(x)
}
