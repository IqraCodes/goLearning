package main

import "fmt"

const (
	a int     = 42
	b float64 = 43.55
	c string  = "Iqra Mohamed"

	d = iota
	f
	g
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

}
