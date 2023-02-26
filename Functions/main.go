package main

import "fmt"

// Function with no parameters nd no return values
func hello() {
	fmt.Println("Hello World")
}

// Function with parameters and unnamed return value
func demo(a, b int) int {
	return a + b
}

// Function with paramaters and named return values
func test(a, b int) (sum, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}

//Variadic function

func check(a int, b ...int) []int {

	fmt.Printf("Type of b is %T \n ", b)
	return b
}

func main() {

	//calling the functions
	hello()

	s := demo(2, 3)
	fmt.Println(s)

	a, b := test(2, 3)
	fmt.Println(a, b)

	f := check(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(f)

	g := check(1)
	fmt.Println(g)

	//Passing a slice to a variadic function
	nums := []int{1, 2, 3, 4, 5}
	h := check(1, nums...)
	fmt.Println(h)

	//Immediately Invoke Function
	square := func(a int) int {
		return a * a
	}(2)

	fmt.Println(square)

	// Anonymous Functions

	var max = func(a, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}

	//using anonymous functions
	res := max(4, 10)
	fmt.Println(res, " is the highest")

	// Anonymous function as an arguement to other function
	result := sno(max(5, 10))
	fmt.Println(result)

	p := practice(4, 5)
	q := p()
	fmt.Println(q)

}
func sno(max int) int {
	return max * max
}

// Returing an anonymous function

func practice(a, b int) func() int {
	fmt.Println("Returning an anonymous function ")
	return func() int {
		return a + b
	}
}
