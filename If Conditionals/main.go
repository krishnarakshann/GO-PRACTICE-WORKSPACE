package main

import "fmt"

func main() {

	//Simple If

	var a = 10
	if a > 5 {
		fmt.Println("A is greater than 10")
	}

	if fmt.Println("Hello World"); true {
		fmt.Printf("Executing if  block\n")
	}

	/*
		if var a1 int = 20;a1>10{
			fmt.Printf("A is greater than 10")
		}
			// Error since var keyword is not supoorted in in statement initialisation
	*/

	if a := 20; a > 10 {
		fmt.Println("value of a which is declared inside if block is ", a)
	}

	fmt.Println("Value of a which is declared outside the if condition is ", a)

	if hello(); true {
		fmt.Println("Executing If block")
	}

}

func hello() {
	fmt.Println("Hello World from hello () function")
}
