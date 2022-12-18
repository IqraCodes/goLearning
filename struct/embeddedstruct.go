//aggregate data type

// How you can take one type and embed it into another type

package main

import "fmt"

type person struct {
	first  string
	middle string
	last   string
	age    int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	p1 := secretAgent{
		
		person: person {
			first:  "Iqra",
			middle: "AbdiRashid",
			last:   "Mohamed",
			age:    27,
		},
		ltk: true
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
