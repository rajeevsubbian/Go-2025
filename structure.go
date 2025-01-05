package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
	Phone Phone
}

type Phone struct {
	AreaCode string
	Prefix   string
	Suffix   string
}

func main() {
	p := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Phone: Phone{
			AreaCode: "123",
			Prefix:   "456",
			Suffix:   "7890",
		},
	}
	fmt.Println(p)
	// q := &Person{"Miss", "Moneypenny", 25}
	// fmt.Println(p.Age)
	// fmt.Println(q.Age)
}
