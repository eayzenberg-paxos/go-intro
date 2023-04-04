package intro

import "fmt"

/*
Embedding in Go is not inheritance, but it can be used to achieve similar
functionality. While inheritance in object-oriented programming allows one class
to inherit fields and methods from another class, embedding in Go allows a
struct to include the fields and methods of another struct, without actually
inheriting from it.

When you embed a struct in another struct, the fields and methods of the
embedded struct become part of the embedding struct. This means that you can
access the fields and methods of the embedded struct using dot notation on the
embedding struct, as if they were its own fields and methods.

Here's an example of embedding in Go that might be used to achieve similar functionality to inheritance:
 */

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println("...")
}

type Dog struct {
	*Animal
}

func (d *Dog) Speak() {
	fmt.Println("Woof!")
}

/*
In this example, we define two struct types: Animal and Dog. The Animal struct
type has a single field Name, and a method Speak() that prints out .... The Dog
struct type embeds an Animal pointer as its only field, which means that it
includes all of the fields and methods of the Animal struct type.

The Dog struct type also has its own Speak() method that overrides the Speak()
method of the Animal struct type. When called on a Dog instance, the Speak()
method of the Dog type will print out Woof! instead of ....

To use the Dog struct type, you can create a new Dog instance and call its Speak() method:
 */

func main() {
	d := &Dog{Animal: &Animal{Name: "Fido"}}
	fmt.Printf("%s says ", d.Name)
	d.Speak()
}

/*
In this example, we create a new Dog instance called d with an embedded Animal
instance. We set the Name field of the Animal instance to "Fido". We then call
the Speak() method of the Dog instance, which will print out Woof!.

Note that while this example achieves similar functionality to inheritance, it
is not actually inheritance. Inheritance is a mechanism of object-oriented
programming that allows one class to inherit fields and methods from another
class, while embedding in Go allows a struct to include the fields and methods
of another struct without inheriting from it.
 */

