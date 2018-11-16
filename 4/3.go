package main

import "fmt"

func main() {
	xx := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(len(xx))
	fmt.Println(xx[:5])
	fmt.Println(xx[5:])
	fmt.Println(xx[2:7])
	fmt.Println(xx[1:6])
}
