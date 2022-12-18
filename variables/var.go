package main

import "fmt"

var a int = 50
var h string = "HELLOOOOOOOO"

var bye string

func main() {
	fmt.Println(h, a)

	bye = `I want to tell you, so much, "one" is "BYEEEEEEEEEE"`

	fmt.Println(bye)
	fmt.Printf("%T", bye)
	fmt.Println(bye)
}
