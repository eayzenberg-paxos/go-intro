package intro

import "encoding/json"

/*
In Go, a deep copy of a data structure can be achieved by using marshalling and
unmarshalling. Marshalling is the process of converting a Go value to a byte
slice, and unmarshalling is the process of converting the byte slice back to a
Go value.

Here's an example of how you can use marshalling and unmarshalling to create a deep copy of a data structure:
 */

// Define a function to create a deep copy of a Person object
func deepCopy(p *Person) *Person {
	// Marshal the Person object to a byte slice
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	// Unmarshal the byte slice to a new Person object
	var copy Person
	err = json.Unmarshal(b, &copy)
	if err != nil {
		panic(err)
	}

	// Return the new Person object
	return &copy
}
