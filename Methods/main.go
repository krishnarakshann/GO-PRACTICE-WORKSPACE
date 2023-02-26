package main

import "fmt"

type Employee struct {
	name string
	age  int
	Address
}
type Address struct {
	Location string
}

// Method with value receiver
func (e Employee) Describe() {
	fmt.Println("Details of Employees are ")
	fmt.Println("Name ", e.name)
	fmt.Println("Age ", e.age)
}

func (e *Employee) ModifyName(newname string) {
	e.name = newname
}

func (a *Address) GetAddress() {
	fmt.Println("Address ", a.Location)
}

// Method on Non Struct Type

type MyInt int

func (m *MyInt) GetValue() {
	fmt.Println("Method on Non struct Type and value of m is ", *m)
}

func main() {

	//create an instance of the struct
	e := Employee{
		name: "krishna",
		age:  23,
		Address: Address{
			Location: "Guntur",
		},
	}
	(&e).Describe()
	fmt.Println("Making changes")
	e.ModifyName("krishna rakshann") //internally it works as (&e).Describe() since it is a pointer receiver
	e.Describe()

	// Method on anonymous nested struct field
	a := Address{
		Location: "Bangalore",
	}
	a.GetAddress()
	e.Address.GetAddress()

	// Method on Non struct Type
	a1 := 2
	b := MyInt(a1)
	(&b).GetValue()

}
