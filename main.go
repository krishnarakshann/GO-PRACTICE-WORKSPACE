package main

import (
	sub "Practice/substract"
	"fmt"
)

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(sub.Substract(3, 4))

}

func add(a, b int) int {
	return a + b
}
