package main

import (
	"fmt"
	"go/day01/calc"
)

func main() {
	fmt.Println("Hello world")

	r1 := calc.Add(10, 20)

	fmt.Println("10 + 20 等于", r1)

	r2 := calc.Sub(10, 20)

	fmt.Println("10 - 20 等于", r2)

	r3 := calc.Multiply(10, 20)

	fmt.Println("10 * 20 等于", r3)

	r4 := calc.Divide(10, 20)

	fmt.Println("10 / 20 等于", r4)
}
