package basic

import (
	"fmt"
	"testing"
)

// In Go, both slices and arrays are used to store ordered collections of elements, but they have some key differences.

func difference() {
	//Size: Arrays have a fixed size that is determined at the time of declaration,
	// while slices can have a variable size that can change dynamically.
	var a [3]int        // array with fixed size 3
	b := []int{1, 2, 3} // slice with variable size
	println(fmt.Sprint(a), fmt.Sprint(b))
}

func copyElements() {
	// Memory allocation: Arrays are allocated on the stack, while slices are allocated on the heap.
	var a [1000000]int        // array with large size, allocated on the stack
	b := make([]int, 1000000) // slice with large size, allocated on the heap
	println(fmt.Sprint(a), fmt.Sprint(b))
}

func mutability() {

	/* Mutability: Arrays are immutable once they are created, meaning that you cannot
	add or remove elements from an array. Slices, on the other hand, are mutable and
	can be modified by adding or removing elements. */

	var a [3]int        // immutable array with fixed size 3
	b := []int{1, 2, 3} // mutable slice with variable size
	b = append(b, 4)    // add an element to the slice
	println(fmt.Sprint(a), fmt.Sprint(b))
}

func modifyArray(arr [3]int) {
	arr[0] = 100
}

func modifySlice(s []int) {
	s[0] = 100
}

func TestSliceVsArray(t *testing.T) {

	/* In this example, we define two functions, modifyArray and modifySlice, that
	take an array and a slice, respectively, and modify the first element.We then
	call these functions with an array a and a slice b, and print the
	results.Since arrays are passed by value, a remains unchanged, while the slice
	b is passed by reference and is modified.

	Overall, arrays are useful for storing fixed-size collections of elements,
	while slices are more flexible and can be used to store variable-size
	collections that can change over time.Slices are also more commonly used in Go
	due to their dynamic size and mutability. */

	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}
	modifyArray(a) // a is unchanged
	modifySlice(b) // b is modified
	fmt.Println(a, b)
}
