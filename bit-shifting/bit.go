package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	// This is bit shifting so it when from 4 to 8 or 100 to 1000
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	// how to use iota in a creative way
	//kilobyte, megabyte and gigabyte in binary

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, kb)
	fmt.Printf("%d\t\t%b\n", gb, kb)

}
