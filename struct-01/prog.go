package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
	gender   string
}

func (p person) fullName() string {
	return fmt.Sprintf("%s %s", p.name, p.lastName)
}

func main() {
	luiz := person{
		name:     "Luiz",
		lastName: "Felipe",
		gender:   "M",
		age:      16,
	}

	fmt.Printf("Hello! My name is %s!\n", luiz.fullName())
}
