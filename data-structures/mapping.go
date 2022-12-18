package main

import "fmt"

func main() {

	m := map[string]int{
		"James": 22,
		"Iqra":  18,
	}

	fmt.Println(m["James"])
	fmt.Println(m["Iqra"])

	fmt.Println(m["Muna"])

	if v, ok := m["Muna"]; ok {
		fmt.Println("This is the if print", v)
	}

	if v, ok := m["Iqra"]; ok {
		fmt.Println("This is the if print", v)
	}

	delete(m, "James")

	fmt.Println(m)
}
