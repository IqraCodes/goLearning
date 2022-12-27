package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a) // gives you the address

	b := &a

	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b)  // gives you the value stores at a address
	fmt.Println(*&a) // dereferences the address then gives you the values inside it

}
