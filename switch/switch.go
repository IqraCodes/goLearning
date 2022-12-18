package main

import "fmt"

// control flow is how a computer reads your programme

func main() {

	n := "Bond"
	switch n {
	case "MoneyBagggggg", "Bond":
		fmt.Println("hi, zeynab")
	case "M":
		fmt.Println("Bond James")
	case "Q":
		fmt.Println("Love that")
	default:
		fmt.Println("hi, iqra")
	}

}
