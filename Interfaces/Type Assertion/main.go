package main

import "fmt"

type demo struct {
}

func checkType(i interface{}) {

	v, ok := i.(int)
	if ok == false {
		fmt.Println("Not an int type ", v)
	}

	a, ok := i.(demo)

	if ok == false {
		fmt.Println("the underlying type is not demo ", a)
	} else {
		fmt.Println("the underlying type is  demo ", a)
	}
}
func main() {
	var a int = 5
	checkType(a)

	d := demo{}
	checkType(d)

	//It is also possible to compare a type to an interface.
	//If we have a type and if that type implements an interface,
	//it is possible to compare this type with the interface it implements.
}
