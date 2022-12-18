// Struct is a data structure that allows you to compose different values pf different types
//aggregate data type

package main

import "fmt"

type person struct {
	first  string
	middle string
	last   string
	age    int
}

func main() {

	p1 := person{
		first:  "Iqra",
		middle: "AbdiRashid",
		last:   "Mohamed",
		age:    27,
	}

	p2 := person{
		first:  "Caliya",
		middle: "AbdiRashid",
		last:   "Mohamed",
		age:    24,
	}

	fmt.Printf("Hi, My name is %v %v. I am %v years old.\n", p1.first, p1.last, p1.age)
	fmt.Printf("Hi, My name is %v %v. I am %v years old.\n", p2.first, p2.last, p2.age)

}
