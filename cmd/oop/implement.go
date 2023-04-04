package intro

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

/*
To implement an interface in Go, you simply need to define a type with methods
that have the same method signatures as the interface. Here's an example of a
Rectangle type that satisfies the Shape interface:
 */

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

/*
In this example, we define a Rectangle type with Width and Height fields, as
well as Area() and Perimeter() methods that calculate the area and perimeter of
the rectangle. Because the Rectangle type has methods with the same method
signatures as the Shape interface, it satisfies the Shape interface.

To use the Rectangle type as a Shape, you can simply create a Rectangle instance
and pass it to any function that expects a Shape. Here's an example:
*/

func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	PrintShapeInfo(r)
}

/*
The statement var _ Shape = (*Rectangle)(nil) in Go is used to check whether a
type satisfies an interface at compile time. In this statement, Shape is the
interface type and Rectangle is the concrete type being checked. The nil value
is used to create a pointer of the Rectangle type. When this statement is
executed, Go will check if the Rectangle type satisfies the Shape interface. If
the Rectangle type does not satisfy the Shape interface, the compiler will throw
an error and the program will not compile. If the Rectangle type does satisfy
the Shape interface, nothing will happen and the program will compile without
errors. Note that the _ before Shape in the statement indicates that we are not
actually using the variable, but we are only interested in checking if the type
satisfies the interface.
 */
var _ Shape = (*Rectangle)(nil)

