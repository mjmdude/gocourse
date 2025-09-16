package main

import (
	"fmt"
	"math"
)

func main() {
	// Variable declaration
	var a, b int = 10, 3
	var c, d float64 = float64(a), float64(b)
	var result int
	var result2 float64
	result = a + b
	fmt.Println("Addition: ", result)
	result = a - b
	fmt.Println("Subtraction: ", result)
	result = a * b
	fmt.Println("Multiplication: ", result)
	result2 = c / d
	fmt.Println("Division: ", result)
	result = a % b
	fmt.Println("Remainder: ", result)

	const p float64 = 22 / 7.0
	fmt.Println(p, "\n", result2)
	// Overflow with signed integer
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)
	maxInt = maxInt + 1
	fmt.Println(maxInt)
	// Overflow with unsigned integer
	var uMaxInt uint64 = 18446744073709551615
	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}
