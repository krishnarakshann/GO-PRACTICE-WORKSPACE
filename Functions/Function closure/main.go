package main

import "fmt"

func fibonacci(a, b int) func() int {

	sum := a + b

	a = b
	b = sum

	return func() int {
		return sum
	}
}

func main() {

	for i := 0; i < 10; i++ {
		fib := fibonacci(i, i+1)
		fmt.Println(fib())
	}
}
