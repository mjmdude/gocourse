package main

import (
	"fmt"
	"slices"
)

func main() {

	// var sliceName []ElementType
	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{9, 8, 7}

	// slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}

	slice2 := a[1:4]

	fmt.Println(slice2)

	slice2 = append(slice2, 6, 7)

	fmt.Println(slice2)

	sliceCopy := make([]int, len(slice2))

	copy(sliceCopy, slice2)

	fmt.Println("Slicecopy", sliceCopy)

	// var nilSlice []int

	for i, v := range slice2 {
		fmt.Println(i, v)
	}

	fmt.Println("Element at index 3 of slice1", slice2[3])

	// slice2[3] = 50

	// fmt.Println("Element at index 3 of slice1", slice2[3])

	if slices.Equal(slice2, sliceCopy) {
		fmt.Println("slice1 is equal to sliceCopy")
	}

	twoD := make([][]int, 3)
	fmt.Println("slice of empty slices:", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}
	fmt.Println(twoD)

	// slice[low:high]
	slice3 := slice2[2:4]
	fmt.Println(slice3)
	fmt.Println("The capacity of slice3 is", cap(slice3))
	fmt.Println("The len of slice3 is", len(slice3))

}
