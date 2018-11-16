package main

import "fmt"

func main() {
	xx := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xx = append(xx, 52)
	fmt.Println(xx)

	xx = append(xx, 53, 54, 55)
	fmt.Println(xx)

	y := []int{56, 57, 58, 59, 60}
	xx = append(xx, y...)
	fmt.Println(xx)
}
