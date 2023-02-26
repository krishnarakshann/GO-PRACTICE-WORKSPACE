package main

import "fmt"

type myString string

func main() {

	// Untyped constants
	const a = 123
	fmt.Println(a)

	// Typed constants -> where we specify the type of constat
	const b = 127

	var b1 int8 = b
	fmt.Println(b1)

	const c = 'e'

	//Getting the default hidden value of untyped  constant
	fmt.Printf("%T is the type and %v is the value \n", b, b)
	fmt.Printf("%T is the type and %v is the value \n", c, c)

	//Finding the hidden type of a unnamed untyped constant
	fmt.Printf("%T is the type and %v is the value \n", 123, 123)

	//strings
	var s1 = "abc\nabc"
	var s2 = `abc\nabc`
	fmt.Printf("%T is the type and %v is the value \n ", s1, s1)
	fmt.Printf("%T is the type and %v is the value \n ", s2, s2)

}
