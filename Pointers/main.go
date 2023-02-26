package main

import "fmt"

func main() {

	a := 2

	b := new(int)
	b = &a

	fmt.Println("A value ", a, " and its adddress  is ", &a)
	fmt.Println("B value  by deferencing it is ", *b, " and its variable adddress  is ", &b, "value inside b without deferencing which is address of a is ", b)

	//making changes in pointer reflects the underlying variable address

	*b = 456
	fmt.Println("A value after modifying the pointer b value ", a)

	//Also * and & can be used together as well. But they will  cancel out each other.

	fmt.Println(*&b)
	fmt.Println(&*b)
	fmt.Println(*&a)
	fmt.Println(b)

	/*
		Pointer to a Pointer
		It is also possible to create a pointer to a pointer in go

		a := 2
		b := &a
		c := &b
		c is a pointer to a pointer here. It stores the address of b which in turn stores the address of a. Double dereferencing using the * operator the can be used to print the value at pointer to pointer.
		Hence  **c will print the value at which 2

	*/

	c := &b
	fmt.Println("Address of variable c ", c, " value inside c ", *c, " which is b value", " double deference the values", **c, "which is val of a")

	/*

		Pointer Arithmetic

		Pointer arithmetic is not possible in golang unlike C language. It raises compilation error.

		e := new(int)
		d := &a
		d = d + e  // Gets an error
	*/

	//zero value

	e := new(int) //not a nil pointer a pointer with no value
	fmt.Println(e, " is the value of e", *e, " is the value inside it ")
	if e == nil {
		fmt.Println("e is nil")
	}

	// nil pointer
	var d *int
	fmt.Println(d, " is the value of d ")

	// Passing a pointer and returning a pointer from a function

	arr := [3]int{1, 23, 45}
	hello(&arr)
	fmt.Println(arr)

	//Passing array as reference to a slice

	demo(arr[:])
	fmt.Println(arr)

}
func hello(arr *[3]int) {
	fmt.Println("Before modifying the value ", *arr)
	(*arr)[0] = arr[0] + 4 // internally it is represented as (*arr)[0]
	arr[1] = arr[1] + 4
	arr[2] = arr[2] + 4
}

func demo(arr []int) {
	fmt.Println("Before modifying the value ", arr)
	arr[0] = arr[0] + 4 // internally it is represented as (*arr)[0]
	arr[1] = arr[1] + 4
	arr[2] = arr[2] + 4
}
