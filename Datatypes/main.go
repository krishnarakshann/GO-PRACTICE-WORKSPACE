package main

import (
	"fmt"
	"math/bits"
	"unsafe"
)

func main() {

	//complex number initialization ways
	n1 := 2 + 3i
	fmt.Println(n1)

	var num complex64 = 2 + 3i
	fmt.Println(num)

	n2 := complex(2, 6)
	fmt.Println(n2)

	/*
		var q1 float32 = 2.3
		var q2 float64 = 4.5
		q3 := complex(q1, q2) // either both q1 and q2 must be of type float32 or float64 when using complex() function
		fmt.Println(q3)   // Error    */

	//using bits package
	sizeOfIntInBits := bits.UintSize
	fmt.Println(sizeOfIntInBits)

	//another way to see the size
	var b = 234
	fmt.Println(unsafe.Sizeof(b))

	var q uint = 0 //takes only 0 or above values cant take negative values
	fmt.Println(q)

	//uintptr
	var name = "rakshann"
	var v1 uintptr
	fmt.Println(v1, " is the initial value of v1")
	fmt.Println(unsafe.Sizeof(v1), " bytes is the size")

	v1 = uintptr(unsafe.Pointer(&name))
	fmt.Println(v1, " is the integer representation of memory location")
	fmt.Println(&name, " is the hexadecimal representation of memory location")

	// Byte --> it is an alias of uint8

	var a byte = 'q'
	fmt.Println(a, " is the value of variable a")
	fmt.Printf("%T is the type and %v is the value of variable a  \n", a, a)

	// if we dont specify the type explicity as byte then it will take rune by default which is an alias of int32

	//rune
	f := 'w'
	fmt.Printf("%T is the type and %v is the value \n ", f, f)
	fmt.Printf("Unicode Point %U \n ", f)
	fmt.Printf("Character %c \n ", f)

}
