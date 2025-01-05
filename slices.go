package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	t := make(([]int), 5)
	fmt.Println(t)
	for i := 0; i < len(s); i++ {
		t[i] = s[i]
	}
	fmt.Println(t)
}
