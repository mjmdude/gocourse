package main

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	elememts []T
}

func (s *Stack[T]) push(element T) {
	s.elememts = append(s.elememts, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elememts) == 0 {
		var zero T
		return zero, false
	}
	element := s.elememts[len(s.elememts)-1]
	s.elememts = s.elememts[:len(s.elememts)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elememts) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elememts) == 0 {
		fmt.Println("The stack is empty.")
		return
	}
	for _, element := range s.elememts {
		fmt.Println(element)
	}
}

func main() {

	// x, y := 1, 2
	// x, y = swap(x, y)
	// fmt.Println(x, y)

	// x1, y1 := "John", "Jane"
	// x1, y1 = swap(x1, y1)
	// fmt.Println(x1, y1)
	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	intStack.pop()
	intStack.printAll()
	fmt.Println("Is Stack empty: ", intStack.isEmpty())
	intStack.pop()
	intStack.printAll()
	fmt.Println("Is Stack empty: ", intStack.isEmpty())
	intStack.pop()
	intStack.printAll()
	fmt.Println("Is Stack empty: ", intStack.isEmpty())

	// For Strings now:

	stringStack := Stack[string]{}
	stringStack.push("John")
	stringStack.push("Fred")
	stringStack.push("Jed")
	stringStack.printAll()
	stringStack.pop()
	stringStack.printAll()
	fmt.Println("Is Stack empty: ", stringStack.isEmpty())
	stringStack.pop()
	stringStack.printAll()
	fmt.Println("Is Stack empty: ", stringStack.isEmpty())
	stringStack.pop()
	stringStack.printAll()
	fmt.Println("Is Stack empty: ", stringStack.isEmpty())
}
