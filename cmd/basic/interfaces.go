package basic

/*
In Go, an interface is a collection of method signatures. An interface defines a
set of behaviors that a type can exhibit, regardless of its implementation. Any
type that implements all the methods of an interface is said to satisfy that
interface, and can be used wherever that interface is expected.
 */

type Shape interface {
	Area() float64
	Perimeter() float64
}

/*
In this example, we define an interface called Shape. The Shape interface has
two method signatures: Area() and Perimeter(). Any type that implements both of
these methods with the correct signatures will satisfy the Shape interface.
 */
