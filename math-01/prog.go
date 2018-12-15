package main

import (
	"fmt"
	"math"
)

func main() {
	num := 64
	sqr := int(math.Sqrt(float64(num)))

	fmt.Printf("The square root of %d is %d!\n", num, sqr)
}
