package main

import "fmt"

func main() {

	// arrays arent really used in Go
	var age [27]int

	fmt.Println(age)

	age[1] = 2

	fmt.Println(age)
	fmt.Println(len(age))

}
