package basic

import (
	"fmt"
	"testing"
)

func TestImmutability(t *testing.T) {

	/* Go is not a pure functional language, so it does not have built-in support for
	immutability like languages such as Haskell or Clojure. However, there are some
	ways to achieve immutability in Go through coding practices and patterns.
	One common approach to achieving immutability in Go is to use immutable data
	structures. This means using data structures that do not allow modification
	after they are created, such as maps or structs with read-only fields. */

	person := Person{Name: "Alice", Age: 30}
	// create a new struct with the same values as person, but with a different name
	newPerson := Person{Name: "Bob", Age: person.Age}
	fmt.Println(person, newPerson)
}
