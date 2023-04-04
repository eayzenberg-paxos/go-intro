package basic

import "fmt"

/*
Error handling in Go is based on the concept of returning errors as values from
functions, rather than throwing exceptions or panicking as in some other
languages. This approach encourages explicit error handling, making code more
robust and easier to reason about.

In Go, errors are represented by the error interface, which has a single method signature:
 */

//	type error interface {
//    	Error() string
//	}

/*
Any type that implements the Error() method with the correct signature is said
to satisfy the error interface, and can be used as an error value.

Functions in Go typically return an error value as their last return value. If
the function completes successfully, it will return nil for the error value. If
an error occurs, the function will return a non-nil error value that describes
the error.

Here's an example of a function that returns an error value:
 */

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

/*
In this example, the Divide() function takes two float64 arguments and returns a
float64 result and an error value. If the second argument b is zero, the
function will return an error value with a message indicating that division by
zero is not allowed. Otherwise, it will return the result of dividing a by b.

To handle errors returned by a function, you can use the if statement to check
if the error value is nil. If the error value is not nil, you can take
appropriate action to handle the error. Here's an example of how you might use
the Divide() function and handle any errors it returns
 */

func main() {
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Result: %f\n", result)
}

/*
In this example, we call the Divide() function with arguments 10 and 2. If the
function returns a non-nil error value, we print an error message and exit the
program. Otherwise, we print the result of the division.

Go also provides the panic() function to terminate a program immediately with an
error message, but it is generally recommended to use error values and explicit
error handling instead of panicking.
 */

