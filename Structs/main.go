package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Employee struct {
	Name string  `json:"n"`
	Sal  float64 `json:"s"`
}

func main() {

	//Pointer to struct

	e := &Employee{}

	(*e).Name = "abc"
	e.Sal = 4.567
	fmt.Println(e)

	f := new(Employee)
	f.Name = "abcd"
	f.Sal = 5678.90
	fmt.Println(f)
	fmt.Printf("Address of p is %p \n", f)
	fmt.Printf("value of pointer %+v\n", *f)
	fmt.Printf("value of pointer %#v\n", *f)

	g := Employee{
		Name: "abcder",
		Sal:  45678.90,
	}
	ejson, err := json.Marshal(g)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(ejson)
	fmt.Println(string(ejson))

	// Anonymous structs - structs with anonymous fields
	type student struct {
		string
		int
		float64
	}

	s := student{
		string:  "abc",
		int:     450,
		float64: 7.8,
	}

	fmt.Println(s)

	// anonymous struct - struct with no name and no field names

	a := struct {
		string
		int
	}{
		string: "krishna",
		int:    789,
	}
	fmt.Printf("struct a is %#v \n", a)

}
