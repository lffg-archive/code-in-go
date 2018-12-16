package main

import "fmt"

// Calculator struct:
type Calculator struct {
	current float64
}

// Sets a new value and returns the new value:
func (c *Calculator) set(value float64) float64 {
	c.current = value
	return value
}

// Sum the passed value to the current value and return the passed value:
func (c *Calculator) add(value float64) float64 {
	c.current = c.current + value
	return value
}

func main() {
	calc := Calculator{current: 1}
	calc.add(2)
	fmt.Println(calc.current)
	calc.set(20)
	calc.add(6)
	fmt.Println(calc.current)
}
