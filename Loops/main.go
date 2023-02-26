package main

import "fmt"

func main() {

	//simple for loop

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}

	// implementing a while loop using for loop

	var a = 0
	for a < 5 {
		fmt.Printf(" %d ", a)
		a++
	}

	//Infinite For loop
	/*

		for  {
			fmt.Println("Hello")
		}

	*/

	//Break Statement

	for i := 0; i < 5; i++ {
		if i > 3 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("Executing out of for loop ")

	// Label break

outer:
	for i := 0; i < 10; i++ {
		fmt.Printf("\n Iteration %d \n", i+1)
	inner:
		for j := 0; j <= i; j++ {

			if j > 5 {
				fmt.Println("Breaking the inner j loop")
				break inner
			}
			fmt.Printf(" i = %d , j= %d \n", i, j)
		}
		if i > 8 {
			fmt.Println("Breaking outer loop ")
			break outer
		}
	}
	fmt.Println("Loop execution completed")

	//Continue

	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			fmt.Printf("%d  is a multiple of 3 \n", i)
			continue
		}
	}

	//Similarly we have label break too

}
