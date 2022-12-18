//package main

import "fmt"

type hotdog int

var x hotdog

var y int

//func main() {
	//s := fmt.Sprintf("%v\n%v\n%v", x, y, z)
	//fmt.Println(s)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
