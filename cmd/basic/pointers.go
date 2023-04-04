package basic

/*
In Go, a pointer is a variable that stores the memory address of another
variable. Pointers allow you to indirectly access and modify the value of a
variable by modifying the value stored in its memory address.

To declare a pointer variable in Go, you use the * symbol followed by the type
of the variable it points to. You can also use the & symbol to get the memory
address of a variable.
*/
func pointer() {
	var x int = 10
	var p *int = &x // p is a pointer to x
	println(p)
}

// built-in pointers: Slices, maps, channels, funcs, errors, interfaces
