package main

import "fmt"

func main() {

	// composite literal data type
	// x := type{values}

	// Slice allows you to group together values of the same type

	x := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(x)

	//x = append(x, 3, 5, 6)

	y := []int{255, 36, 4889, 52, 4, 68585, 78382}

	y = append(y, 5, x...)

	fmt.Println(y)

}
