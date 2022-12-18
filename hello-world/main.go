package main

import "fmt"

func main() {
	fmt.Println("Hello, how are you")

	foo()

	fmt.Println("Something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {

	fmt.Println("I'm in food")
}

func bar() {

	fmt.Println("EXITEDDDDDDDDDD")

	x := 20

	fmt.Println(x)

	x = 10

	fmt.Println(x)
}
