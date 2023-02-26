package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l *lion) breathe() {
	fmt.Println(l.age, " years lion will brathe")
}

func (l *lion) walk() {
	fmt.Println("It walks")
}
func main() {
	var a animal
	l := lion{}
	a = &lion{age: 20}
	a.walk()
	a.breathe()

	Describe(&l)

}

func Describe(a animal) {
	fmt.Printf("%T and %v of interfce \n", a, a)
}
