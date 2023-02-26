package main

import "fmt"

var ch rune

func main() {

	switch initialise(); ch {

	case 'a':
		fmt.Println("Its case a")
		break
	default:
		fmt.Println("None of the case matches")
	}

	fmt.Println("Switch 2")
	var ch rune
	switch ch = 'r'; ch {
	case 'a', 'r':
		fmt.Println("Its case a or r")
		break
	default:
		fmt.Println("none of the case matches ")
	}

	fmt.Println("Switch 3 only having the expression")
	fmt.Println("ch is ", ch)
	switch ch {
	case f1(), f2():
		fmt.Println("Its a character")
	default:
		fmt.Println("None of the case matches ")
	}

	fmt.Println("Both of the expression and statement are absent")
	age := 45
	switch {
	case age < 18:
		fmt.Println("Kid")
	case age >= 18 && age < 40:
		fmt.Println("Young")
	default:
		fmt.Println("Old")
	}

	fmt.Println("Only  statement  is presnt")

	switch age := 45; {
	case age < 18:
		fmt.Println("Kid")
	case age >= 18 && age < 40:
		fmt.Println("Young")
	default:
		fmt.Println("Old")
	}

	// Two case statement cannot have the same constant

	// Fallthroughs
	exp := 10
	switch {
	case exp > 0:
		fmt.Println("Exp is greater than 0")
		fallthrough
	case exp < 0:
		fmt.Println("Exp is less then 0")
	}

}

func initialise() {
	ch = 'a'
}

func f1() rune {
	return 'a'
}

func f2() rune {
	return 'r'
}
