package main

import "fmt"

var (
	name = "rakshann"
)

func main() {

	// Five types of variable declarations

	//single variable declaration without initial value

	var a int
	fmt.Println(a)

	//single variable declaration with intial value
	var b int = 10
	fmt.Println(b)

	// multiple variable declarations without intial values
	var c, d, e int
	fmt.Println(c, d, e)

	// Multiple variable declarations with initial values
	var a1, a2, a3 int = 10, 20, 30
	fmt.Println(a1, a2, a3)

	//declaring variables of different types

	var (
		b1 int
		b2 string  = "rakshann"
		b3 float64 = 9.89
	)
	fmt.Println(b1, b2, b3)

	// variable declaration with no Type or Type inference
	var name = "krishna"
	fmt.Println(name)
	/*The GO compiler based upon the value assigned to the variable will figure out the type. So if the variabl
	e has an initial value, then the type can be omitted.  This is also called Type Inference.
	Below is the format for such declaration
	*/

	//Short variable declaration --> atleast one of the variable on the left hand side of := must be unique or new when using this
	age := 23
	fmt.Println(age, " is the age")

	// The local varibale scope is greater than global variables the local variable shadows the outer variable
	name = "krishna"
	fmt.Println("Inner function scope is ", name)

}
