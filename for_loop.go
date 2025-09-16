package main

import "fmt"

func main() {
	// Simple iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	numbers := []int{1, 2, 3, 4, 5, 6}
	for i, v := range numbers {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	// for break.  break gets you out of the for loop
	// for continue. continue skips the remaining code in the for block and goes to the next iteration
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue //continue the loop but skip everthing below
		}
		fmt.Println("Odd Number: ", i)
		if i == 5 {
			break // break out of the forloop block
		}
	}
	rows := 5
	//Outer loop
	for i := 1; i <= rows; i++ {
		// inner loop for spaces before stars
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		//inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // Move to the next line
	}
	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("We have a lift off.")
}
