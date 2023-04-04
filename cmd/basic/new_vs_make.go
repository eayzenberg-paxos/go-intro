package basic

/*
In Go, make and new are two built-in functions used to create new values of different types.

new is used to allocate memory for a new value of a given type, and returns a
pointer to the newly allocated memory. The new function takes a single argument,
which is the type of the value to be created.

Here's an example of using new to allocate memory for a new integer value:
 */
func _new() {
	var p *int = new(int)
	*p = 10
}
/*
In this example, we use new to allocate memory for a new integer value, and
store the memory address of the newly allocated memory in a pointer variable p.
We then set the value of the newly allocated memory to 10.

On the other hand, make is used to create a new slice, map, or channel. The make
function takes two or three arguments: the type of the value to be created, the
length of the slice, map, or channel (if applicable), and the capacity of the
slice (if applicable).

Here's an example of using make to create a new slice:
 */
func _make() {
	var s []int = make([]int, 5)
	s[0] = 10
}
