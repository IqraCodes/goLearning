package main

import "fmt"

//var x string = `Hello, this is a raw string literal so what can you "SEEEEEEEEEEE" lalallaal`

const (
	year1 = 2018 + iota
	year2
	year3
	year4
)

func main() {

	//x := 42
	//fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	//y := x << 1
	//fmt.Printf("%d\t%b\t%#x", y, y, y)

	fmt.Println(year1, year2, year3, year4)

}
