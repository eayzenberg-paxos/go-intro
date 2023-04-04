package basic

import "fmt"

// map[keyType]valueType

var m map[int]string // This creates an empty map that can store integers as keys and strings as values.

func putValue() {
	m[1] = "value"
}

func makingAmap() {
	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
}

func retrieveFomMap() {
	value := m[1]; println(value)
	fmt.Println(m[1]) // Output: "one"
	fmt.Println(m[4]) // Output: "" (empty string)

	value, ok := m[5] // here ok is false
	println(ok)
}

func deleteFromMap() {
	delete(m, 2)   // built-in func
	fmt.Println(m) // Output: map[1:one 3:three]
}