package main

import "fmt"

func main() {
	s := "My String"
	sptr := &s
	fmt.Println(s)
	fmt.Println(sptr)
	fmt.Println(*sptr)
}
