package intro

import "fmt"

type Person struct {
	Name string
	Age  int
}

/*
In Go, a method is a function associated with a type. Methods are similar to
functions, but are called on a specific instance of a type, and can access and
modify the state of that instance.

In Go, methods are defined using the func keyword, followed by the receiver type
and a method name. The receiver type is the type that the method is associated
with, and is specified using a type name (either a named or an anonymous type).

Here's an example of a method definition in Go: */

func (p Person) SayHello() { // here p is called 'receiver' and works as 'this'
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "John", Age: 42}
	p.SayHello()
}

/*
In Go, methods with pointer receivers are a way to modify the state of a struct.
When a method is defined with a pointer receiver, the method receives a pointer
to the struct instance instead of a copy of it. This means that any changes made
to the struct inside the method are reflected in the original struct instance.
 */

func (p *Person) IncreaseAge() {
	p.Age++
}

/*
In this example, we define a Person struct and a method on that struct called
IncreaseAge(). The method has a pointer receiver, which means that it takes a
pointer to a Person struct instance as its receiver.

The IncreaseAge() method increments the Age field of the Person struct instance that it is called on.

To use the IncreaseAge() method, we create a Person struct instance and call the method on it:
 */

func main2() {
	p := &Person{Name: "Alice", Age: 30}
	fmt.Printf("Before: %v\n", p)
	p.IncreaseAge()
	fmt.Printf("After: %v\n", p)

	//In this example, we create a new Person struct instance called p with the Name
	//field set to "Alice" and the Age field set to 30. We then call the
	//IncreaseAge() method on the p instance, which increments its Age field.
	//Finally, we print out the p instance before and after calling the method.
	//
	//Note that we created p as a pointer to a Person struct instance using the &
	//operator. This is necessary because the IncreaseAge() method has a pointer
	//receiver, which means that it can only be called on a pointer to a Person
	//struct instance. If we had created p as a value of type Person instead, we
	//would not be able to call the IncreaseAge() method on it.
}


