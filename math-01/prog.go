package main

import (
	"fmt"
	"math"
)

func print(num int, sqr int) {
	fmt.Printf("The square root of %d is %d!", num, sqr)
}

func main() {
	num := 64
	sqr := math.Sqrt(float64(num))

	print(num, int(sqr))
}
