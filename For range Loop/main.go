package main

import "fmt"

func main() {

	//creating an array
	arr := [10]int{1, 2, 3}
	fmt.Println(arr)

	//using a for range loop

	for index, value := range arr {
		fmt.Printf("%d  index has value %d \n ", index, value)
	}

	//having only the index and ignoring value
	for index := range arr {
		fmt.Printf("%d ", index)
		fmt.Printf("\n")
	}

	// having only the value and ignoring the index
	for _, value := range arr {
		fmt.Printf("%d -> ", value)
		fmt.Printf("\n")
	}

	//ignoring both index and value
	k := 0
	for range arr {
		fmt.Printf("%d -> %d \n ", k, arr[k])
		k++
	}

	//Iterating over the strings
	//len(string) -> returns the number of bytes in the string
	// for range  -> returns the number of unicode points in string

	var name string = "a£b"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c is the character at index %d \n ", name[i], i)
	}

	for index, value := range name {
		fmt.Printf("%d index has character %c \n", index, value)
	}
}
