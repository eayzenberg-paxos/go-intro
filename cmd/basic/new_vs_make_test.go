package basic

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {

	/* In Go, make and new are two built-in functions used to create new values of different types.
	new is used to allocate memory for a new value of a given type, and returns a
	pointer to the newly allocated memory. The new function takes a single argument,
	which is the type of the value to be created.
	In this example, we use new to allocate memory for a new integer value, and
	store the memory address of the newly allocated memory in a pointer variable p.
	We then set the value of the newly allocated memory to 10. */

	var p *int = new(int)
	*p = 10

	println(fmt.Sprintf("Pointer: %v, Value: %v", p, *p))
	// outputs:
	// Pointer: 0xc00011a190, Value: 10
}

func TestMake(t *testing.T) {

	/* On the other hand, make is used to create a new slice, map, or channel. The make
	function takes two or three arguments: the type of the value to be created, the
	length of the slice, map, or channel (if applicable), and the capacity of the
	slice (if applicable). */

	var s []int = make([]int, 5)
	s[0] = 10

	println(fmt.Sprintf("Value: %v", s))
	// outputs:
	// Value: [10 0 0 0 0]
}

func TestNewVsMake(t *testing.T) {

	var s *[]int = new([]int)
	*s = make([]int, 5)
	(*s)[0] = 10

	println(fmt.Sprintf("Pointer: %v, Value: %v", s, *s))
	// outputs:
	// Pointer: &[10 0 0 0 0], Value: [10 0 0 0 0]
}
