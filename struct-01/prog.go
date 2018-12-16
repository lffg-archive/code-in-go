package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
	gender   string
}

func main() {
	luiz := person{
		name:     "Luiz",
		lastName: "Felipe",
		gender:   "M",
		age:      16,
	}

	fmt.Printf("Hello! My name is %s %s!\n", luiz.name, luiz.lastName)
}
