package main

import "fmt"

func main() {
	// In the main function, create a slice of ints from 0 through 10
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Iterate through the slice with a for loop.  For each element, check to see if the number is even or odd.
	for i := range ints {
		// If the value is even, print out "even".  If it is odd, print out "odd"
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}

	}

}
