package basic

// Creating aliases for existing types:

type MyInt int       // Create an alias for the int type
type MyString string // Create an alias for the string type

// Creating new named types based on existing types

type Celsius float64    // New named type based on float64
type Fahrenheit float64 // New named type based on float64

// Creating new named struct types

type Person struct {
	Name    string
	Age     int
	Address string
}
