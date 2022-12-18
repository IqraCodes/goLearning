package main

import "fmt"

func main() {

	//x := 42
	//fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	//y := x << 1
	//fmt.Printf("%d\t%b\t%#x", y, y, y)

	for i := 10; i < 100; i++ {
		fmt.Printf("When %v is divided by 4 the remainder aka modulus is %v\n", i, i%4)
	}

}
