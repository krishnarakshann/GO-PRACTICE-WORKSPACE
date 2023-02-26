package main

import "fmt"

type person struct {
	salary int
	exp    float64
}

func main() {

	//ways to create a map

	var m = map[int]int{
		1: 10,
		2: 20,
	}

	if m == nil {
		fmt.Println("Its a nil map")
	}

	fmt.Println("Map values ", m)

	//using make function

	m = make(map[int]int)
	m[1] = 100
	m[3] = 456
	fmt.Println(m)

	// map where key is a struct

	var m1 = map[person]string{}
	m1[person{
		salary: 1000,
		exp:    8.5,
	}] = "Krishna prasad"

	fmt.Println(m1)

	//Map operations

	var m2 = make(map[int]string)

	//adding a key
	m2[1] = "krishna rakshann"
	fmt.Println(m2)

	//updating an existing key ->it will create a new key-value pair it doesnt exist or it will update the existing key with new value

	m2[1] = "Vijay Rakshann"
	fmt.Println(m2)

	//Fetching the existing value of key
	fmt.Println("Value of key 1 ", m2[1])

	//Trying to fetch a key which even doesn't return's the default value of value type
	fmt.Println(m2[34], " is the value")

	//Delete the key
	delete(m2, 1)
	fmt.Println("After deleting the key ", m2)

	delete(m2, 89)
	fmt.Println("after delete the key which even doesnt exist ", m2)

	//Chekcing if a key exists or not

	v, ok := m2[89]
	if ok == false {
		fmt.Println("Key Doesnt Exists")
	} else {
		fmt.Println("Value of Key is ", v)
	}

	//creating a map where key is an array

	m3 := make(map[[5]int]string)
	m3[[5]int{1, 2, 3}] = "krishna rakshann"
	fmt.Println(m3, " is the value ")

	//map where value is an array

	m4 := make(map[string][5]int)
	m4["krishna rakshann"] = [5]int{1, 2, 3, 4, 5}
	fmt.Println(m4, "is the value of m4")

	//Assign one map to other

	var name = map[int]string{}
	name[1] = "krishna"
	name[2] = "rakshann"
	name[3] = "vijay"
	fmt.Println("Final values of names ", name)

	var dup map[int]string = name
	fmt.Println("Dup value is ", dup)

	//Map is a reference when it is assigned to other variable making chanes in one map reflects in another

	dup[4] = "stunt kid"
	fmt.Println("After modifying changes to map ", dup, name)

	delete(dup, 2)
	fmt.Println("After modifying changes to map ", dup, name)

	// Iteration over a map

	for key, value := range name {
		fmt.Println("Key -> ", key, " value is -> ", value)
	}

	// the order of output keeps on changing in maps you cant predict it

	// using only for loop (not suggested to use)
	for i := 0; i <= len(name); i++ {
		fmt.Print(name[i], "***")
	}

	// fetching only keys
	keys := []int{}
	for key, _ := range name {
		keys = append(keys, key)
	}
	fmt.Println("\n List of keys are ", keys)

}
