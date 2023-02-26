package main

import "fmt"

func main() {

	//Ways of Array declaration

	var a [5]int
	fmt.Println(a)

	var b = [5]int{1, 2, 3}
	fmt.Println(b)

	//array of inifinite length ... is an instruction to compiler to find the length of the array
	var c = [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	var d = [...]int{0: 1, 10: 12}
	fmt.Println(d)

	//size of an array is part of the type of arrays of different length will be treated as array of differenr type

	var e = [4]int{1, 2}
	var f = [4]int{2, 3}

	if e == f {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//arrays are value types -- when they are assigned to a variable the copy of array is actually assigned

	g := f
	fmt.Println(g)
	g[1] = 1209
	fmt.Println("After changes the value of g is ", g)
	fmt.Println("After changes value of f is ", f)

	h := &g
	fmt.Println("H value is ", h)
	g[0] = 1098
	fmt.Println("H value is after modifying the underlying g ", h)

	//another way of declaring and using arrays is using short-hand assignment

	a1 := [5]int{}
	fmt.Println(a1)

	a2 := [5]int{1, 2}
	fmt.Println(a2)

	a5 := a2[0:3]
	fmt.Println("Value of a5 is ", a5)
	fmt.Println("Length of a5 is ", len(a5))

	a3 := [3]int{0: 1, 1: 2}
	fmt.Println(a3)

	a4 := [...]int{}
	fmt.Println("Value of a4 is ", a4)

	//Multidimensional Arrays

	var m = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Multidimensional array is ", m)
	fmt.Println("Length of the Multidimensional array is ", len(m))
	fmt.Println(m[0])

	//iterating over the array - we can use for loop or for-range loop

	for _, sample := range m {
		for index, val := range sample {
			fmt.Println("Index ", index, " had value ", val)
		}
	}

}
