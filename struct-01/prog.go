package main

import "fmt"

// Creates a "person" struct:
type person struct {
	name     string
	lastName string
	age      uint8
	gender   string
}

// Attaches a method to the person struct:
func (p person) fullName() string {
	return fmt.Sprintf("%s %s", p.name, p.lastName)
}

func main() {
	// Creates a new instance of the person struct:
	luiz := person{
		name:     "Luiz",
		lastName: "Felipe",
		gender:   "M",
		age:      16,
	}

	fmt.Printf("Hello! My name is %s!\n", luiz.fullName())
}
