package main

import "fmt"

func main() {

	// composite literal data type
	// x := type{values}

	// Slice allows you to group together values of the same type

	x := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Println(i, v)
	}

	// Slicing a Slice

	y := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(y)

	fmt.Println(y[1:4])

	v := len(x)
	for y {

		fmt.Println(y, v)
	}

}
