package main

import "fmt"

func main() {

	var a = []int{} //empty slice
	fmt.Println(a)

	if a == nil {
		fmt.Println("Its a nil slice")
	}

	var b []int //nil slice
	fmt.Println(b)

	if b == nil {
		fmt.Println("Its a nil slice")
	}

	var c = []int{1, 2, 3, 4, 5, 6, 7, 9, 100: 20}
	fmt.Println(c)

	//creating a slice from another slice

	var d = c[10:20] //ending index is not inclusive
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d), " is the cpacity")

	c[3] = 1230
	fmt.Println(d)

	e := c[25:]
	fmt.Println(e)
	fmt.Println(len(e))
	fmt.Println(cap(e), " is the cpacity")

	f := c[:5]
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f), " is the cpacity")

	g := c[:] // copying the entire slice
	fmt.Println(g)
	fmt.Println(len(g))
	fmt.Println(cap(g), " is the cpacity")

	// slice creation using make function

	var a1 = make([]int, 10)
	fmt.Println(a1)

	// creation of slice  or array using new function

	var nums = new([5]int)
	fmt.Println(*nums)
	nums[0] = 100
	fmt.Println(*nums)

	var v1 = []int{1, 2, 3, 4, 5, 6}
	var v2 = []int{4, 5}
	c2 := copy(v2, v1) // returns the number of elements it copied
	fmt.Println(c2)    //2

	//Multi-dimensional slice

	var v4 = [][]int{
		{1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10},
	}

	for _, val := range v4 {
		for _, v := range val {
			fmt.Print(v, "-")
		}
	}

	stunt := make([]int, 0)
	if stunt == nil {
		fmt.Println("it is a nil slice ")
	} else {
		fmt.Println("\n Not a nil slice ")
	}

	/*
		slice can be compared only with nil value
	*/

}
