package add

import "fmt"

func Addition(a, b int) int {
	sum = a + b
	return sum
}

var sum int

func init() {
	fmt.Println("In package add intialising the sum value")
	sum = 0

}

// can have more than one init funciton for a file
func init() {
	fmt.Println("In second init() function")
}
