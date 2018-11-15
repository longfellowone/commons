package main

import "fmt"

func main() {
	x := 3
	switch {
	case x == 10:
		fmt.Println("x = 10")
	case x == 5:
		fmt.Println("x = 5")
	case x == 0:
		fmt.Println("x = 0")
	default:
		fmt.Println("x is not 0,5 or 10")
	}
}
