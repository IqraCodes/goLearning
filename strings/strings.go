package main

import "fmt"

func main() {
	s := "Hello byeeeeeee bisshhhhhhh"
	fmt.Println(s)
	fmt.Println("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println("%T\n", bs)
}
