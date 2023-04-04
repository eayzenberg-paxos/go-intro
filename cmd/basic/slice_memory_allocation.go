package basic

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1
	s2 = append(s2, 4)
	fmt.Println(s1, s2)
}

/*
In this example, we create a slice s1 with three elements, and then create a
second slice s2 that refers to the same underlying array as s1. We then append
the element 4 to s2, causing Go to allocate a new array with larger capacity,
copy the existing elements to the new array, and append the new element to the
end of the new array.
 */
